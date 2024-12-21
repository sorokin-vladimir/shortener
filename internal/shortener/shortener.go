package shortener

import (
	"errors"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sorokin-vladimir/shortener/internal/database"
	"github.com/sorokin-vladimir/shortener/internal/utils"
)

type ReturnType struct {
	Short          string
	expiry         time.Duration
	remainingQuota int
	Err            error
	rateLimitReset time.Duration
	_              string
}

func Shortener(
	url string,
	IP string,
	ID int64,
	customExpiry int,
	customShort string,
) ReturnType {
	log.SetPrefix("Shortener | ")
	db_limits := database.CreateClient(database.DB_LIMITS)
	defer db_limits.Close()

	var userID string
	if IP != "" {
		userID = IP
	} else {
		userID = strconv.FormatInt(ID, 10)
	}
	log.Printf("Try to short URL: %s | User: %s", url, userID)

	val, err := db_limits.Get(database.Ctx, userID).Result()
	limit, _ := db_limits.TTL(database.Ctx, userID).Result()

	apiQuota, quotaErr := strconv.Atoi(os.Getenv("API_QUOTA"))
	if quotaErr != nil {
		log.Printf("Quota err: %e", quotaErr)
		return ReturnType{
			Short:          "",
			expiry:         0,
			remainingQuota: 0,
			Err:            errors.New(utils.INTERNAL_ERROR),
			rateLimitReset: 0,
		}
	}

	if err == redis.Nil {
		_ = db_limits.Set(database.Ctx, userID, apiQuota, 30*60*time.Second).Err()
	} else if err == nil {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			log.Printf("Rate limit exceeded for user %s", userID)
			return ReturnType{
				Short:          "",
				expiry:         0,
				remainingQuota: 0,
				Err:            errors.New(utils.RATE_LIMIT_EXCEEDED),
				rateLimitReset: limit / time.Nanosecond / time.Minute,
			}
		}
	}

	if len(url) < 4 {
		log.Printf("Too short origin URL. Length is %d", len(url))
		return ReturnType{
			Short:          "",
			expiry:         0,
			remainingQuota: 0,
			Err:            errors.New(utils.INVALID_URL_TOO_SHORT),
			rateLimitReset: 0,
		}
	}

	enforcedURL := utils.EnforceHTTP(url)
	if !utils.IsURL(enforcedURL) {
		log.Printf("Input string does not look like URL: %s. User: %s", enforcedURL, userID)
		return ReturnType{
			Short:          "",
			expiry:         0,
			remainingQuota: 0,
			Err:            errors.New(utils.INVALID_URL),
			rateLimitReset: 0,
		}
	}

	if utils.IsLocalhost(enforcedURL) {
		log.Printf("Tried to shorten 'localhost'. User: %s", userID)
		return ReturnType{
			Short:          "",
			expiry:         0,
			remainingQuota: 0,
			Err:            errors.New(utils.INVALID_URL_LOCALHOST),
			rateLimitReset: 0,
		}
	}

	var id string
	if customShort == "" {
		id = encode(rand.Uint64())
	} else {
		log.Printf("Custom short is provided: %s. User: %s", customShort, userID)
		id = customShort
	}

	db_shorts := database.CreateClient(database.DB_SHORT)
	defer db_shorts.Close()

	val, _ = db_shorts.Get(database.Ctx, id).Result()

	if val != "" {
		log.Printf("The short '%s' is already exist. User: %s", id, userID)
		return ReturnType{
			Short:          "",
			expiry:         0,
			remainingQuota: 0,
			Err:            errors.New(utils.SHORT_EXISTS),
			rateLimitReset: 0,
		}
	}

	var expiry int
	if customExpiry == 0 {
		expiry, _ = strconv.Atoi(os.Getenv("EXPIRY_HOURS"))
	} else {
		log.Printf("Custom expiry is provided: %d. User: %s", customExpiry, userID)
		expiry = customExpiry
	}
	expiryTimeDuration := time.Duration(expiry * 3600 * int(time.Second))

	err = db_shorts.Set(database.Ctx, id, enforcedURL, expiryTimeDuration).Err()

	if err != nil {
		log.Fatalf("Set short in DB err: %e", err)
		return ReturnType{
			Short:          "",
			expiry:         0,
			remainingQuota: 0,
			Err:            errors.New(utils.INTERNAL_ERROR),
			rateLimitReset: 0,
		}
	}

	remainingQuota, _ := db_limits.Decr(database.Ctx, userID).Result()

	var domain string
	if os.Getenv("DOMAIN") == "localhost" {
		domain = os.Getenv("DOMAIN") + ":" + os.Getenv("WEB_PORT") + "/"
	} else {
		domain = os.Getenv("DOMAIN") + "/"
	}

	return ReturnType{
		Short:          domain + id,
		expiry:         expiryTimeDuration,
		remainingQuota: int(remainingQuota),
		Err:            nil,
		rateLimitReset: 0,
	}
}
