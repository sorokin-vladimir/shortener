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

func TestDecode(t *testing.T) {

	for _, tt := range TestData {
		t.Run(fmt.Sprintf("Decode %s", tt.encodedStr), func(t *testing.T) {
			result, err := decode(tt.encodedStr)

			if err != nil {
				t.Error(err)
			}

			if result != tt.number {
				t.Errorf("Decode(%s) = %d; expected %d", tt.encodedStr, result, tt.number)
			}
		})
	}

	const encodedStrWithWrongSymbol = "abc_e"
	t.Run(fmt.Sprintf("Decode %s", encodedStrWithWrongSymbol), func(t *testing.T) {
		result, err := decode(encodedStrWithWrongSymbol)

		if err != nil && err.Error() == "wrong symbol" {
			t.Logf("Decode(%s): expected error", encodedStrWithWrongSymbol)
			return
		}

		t.Errorf("Decode(%s) = %d; expected error", encodedStrWithWrongSymbol, result)
	})
}

var TestData = []TestCase{
	{number: 0, encodedStr: ""},
	{number: 1, encodedStr: "b"},
	{number: 998, encodedStr: "1u"},
	{number: 999, encodedStr: "2u"},
	{number: 15000000, encodedStr: "jkAGb"},
	{number: 100000000, encodedStr: "tWAkm"},
	{number: 10474085661195148593, encodedStr: "VJbdBW332GN"},
	{number: 10998849939277099907, encodedStr: "PQzMntdPVrQ"},
	{number: 3366573584661143826, encodedStr: "vtHkZXQyKfn"},
	{number: 13989813722318721924, encodedStr: "FE16hUaB9g1"},
	{number: 7819934376890328252, encodedStr: "Fx4FtVTFZXC"},
	{number: 2312079729833425620, encodedStr: "pTMMd7vH4Nh"},
	{number: 5235261988083167128, encodedStr: "KMV-bQuGMru"},
	{number: 11660458361119275188, encodedStr: "pPFEe88r5BS"},
	{number: 7409133607407091621, encodedStr: "fQh2PQme9AB"},
	{number: 3277694648817963827, encodedStr: "weVvDwhBkZm"},
	{number: 18176449534139642514, encodedStr: "NcaAuABX86db"},
	{number: 4507735593119348904, encodedStr: "3bMz2YBvs3r"},
	{number: 457861098005776618, encodedStr: "Nhy8pNXaHFb"},
	{number: 14796952220313002321, encodedStr: "brVptK2W-X3"},
	{number: 1103397045671477896, encodedStr: "Bvcq8dNDUPd"},
	{number: 13124883478096200606, encodedStr: "9wR6RPwJssX"},
	{number: 7788278420941980598, encodedStr: "rwzAVNSn9RC"},
	{number: 7571599580588466967, encodedStr: "26Ab4ah488B"},
	{number: 15331931484882232747, encodedStr: "6CUfM1KhXH5"},
	{number: 6867411904295255519, encodedStr: "JUxjf7G-7Pz"},
	{number: 18446744073709551615, encodedStr: "sDC7JC3C3Zeb"},
}
