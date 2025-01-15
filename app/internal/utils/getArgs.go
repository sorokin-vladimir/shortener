package utils

import (
	"log"
	"strconv"
	"strings"
)

func GetArgs(str string) (url string, short string, expiry int) {
	splitedStr := strings.Split(str, " ")

	for i, subStr := range splitedStr {
		switch i {
		case 0:
			url = subStr
		case 1:
			short = subStr
		case 2:
			exp, err := strconv.Atoi(subStr)
			if err != nil {
				log.Println("Could not convert arg into int:", subStr)
				expiry = 0
			} else {
				expiry = exp
			}
		}
	}

	return url, short, expiry
}
