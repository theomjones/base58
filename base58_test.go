package base58

import "testing"

type decodeTable struct {
	want int
	give string
}

type encodeTable struct {
	want string
	give int
}

func TestEncode(t *testing.T) {
	ts := []encodeTable{
		{"2j", 100},
		{"Gw", 924},
		{"44R", 10290},
		{"5Lxy", 847610},
	}

	for _, v := range ts {
		got := Encode(v.give)
		if got != v.want {
			t.Errorf("Got %v but wanted %v", got, v.want)
		}
	}
}

func TestDecode(t *testing.T) {

	ts := []decodeTable{
		{100, "2j"},
		{90218, "TpV"},
		{32467264757, "rU2Dfr"},
		{1, "2"},
		{12, "D"},
	}

	for _, v := range ts {
		got := Decode(v.give)
		if got != v.want {
			t.Errorf("Got %v but wanted %v", got, v.want)
		}
	}
}
