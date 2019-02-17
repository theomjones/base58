package base58

import (
	"strings"
)

func createString() []string {
	ss := strings.Split("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz", "")
	return ss
}

// Encode encodes an unsigned integer to base58.
func Encode(n uint) string {

	ss := createString()

	var rr []uint

	i := n

	for i > 0 {
		r := i % 58
		rr = append(rr, r)
		i = i / 58
	}

	var sl []string

	for _, x := range rr {
		sl = append([]string{ss[x]}, sl...)
	}

	return strings.Join(sl, "")
}
