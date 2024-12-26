package shortener

import (
	"fmt"
	"testing"
)

type TestCase struct {
	number     uint64
	encodedStr string
}

func TestEncode(t *testing.T) {
	for _, tt := range TestData {
		t.Run(fmt.Sprintf("Encode %d", tt.number), func(t *testing.T) {
			result := encode(tt.number)
			if result != tt.encodedStr {
				t.Errorf("Encode(%d) = %s; expected %s", tt.number, result, tt.encodedStr)
			}
		})
	}
}

var TestData = []TestCase{
	{number: 0, encodedStr: ""},
	{number: 1, encodedStr: "b"},
	{number: 998, encodedStr: "1u"},
	{number: 999, encodedStr: "2u"},
	{number: 15000000, encodedStr: "jkAGb"},
	{number: 100000000, encodedStr: "tWAkm"},
	{number: 10474085661195148593, encodedStr: "VJbd"},
	{number: 10998849939277099907, encodedStr: "PQzM"},
	{number: 3366573584661143826, encodedStr: "vtHk"},
	{number: 13989813722318721924, encodedStr: "FE16"},
	{number: 7819934376890328252, encodedStr: "Fx4F"},
	{number: 2312079729833425620, encodedStr: "pTMM"},
	{number: 5235261988083167128, encodedStr: "KMV-"},
	{number: 11660458361119275188, encodedStr: "pPFE"},
	{number: 7409133607407091621, encodedStr: "fQh2"},
	{number: 3277694648817963827, encodedStr: "weVv"},
	{number: 18176449534139642514, encodedStr: "NcaAb"},
	{number: 4507735593119348904, encodedStr: "3bMz"},
	{number: 457861098005776618, encodedStr: "Nhy8"},
	{number: 14796952220313002321, encodedStr: "brVp"},
	{number: 1103397045671477896, encodedStr: "Bvcq"},
	{number: 13124883478096200606, encodedStr: "9wR6"},
	{number: 7788278420941980598, encodedStr: "rwzA"},
	{number: 7571599580588466967, encodedStr: "26Ab"},
	{number: 15331931484882232747, encodedStr: "6CUf"},
	{number: 6867411904295255519, encodedStr: "JUxj"},
	{number: 18446744073709551615, encodedStr: "sDC7b"},
}
