package base58_test

import (
	"testing"

	"github.com/theomjones/base58"
)

func TestEncode(t *testing.T) {
	tt := []struct {
		want string
		give int
	}{
		{"2j", 100},
		{"Gw", 924},
		{"44R", 10290},
		{"5Lxy", 847610},
	}

	for _, v := range tt {
		got := base58.Encode(v.give)
		if got != v.want {
			t.Errorf("Got %v but wanted %v", got, v.want)
		}
	}
}

func TestDecode(t *testing.T) {

	tt := []struct {
		want int
		give string
	}{
		{100, "2j"},
		{90218, "TpV"},
		{32467264757, "rU2Dfr"},
		{1, "2"},
		{12, "D"},
	}

	for _, v := range tt {
		got := base58.Decode(v.give)
		if got != v.want {
			t.Errorf("Got %v but wanted %v", got, v.want)
		}
	}
}
