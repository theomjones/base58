package base58

import (
	"math"
	"strings"
)

const an = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func createString() []string {
	ss := strings.Split(an, "")
	return ss
}

// Encode encodes an unsigned integer to base58.
func Encode(n int) string {

	ss := createString()

	var rr []int

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

func Decode(s string) int {
	ss := strings.Split(s, "")
	var ii []int

	var iiRev []int

	for _, char := range ss {
		ii = append(ii, strings.Index(an, char))
	}

	for _, num := range ii {
		iiRev = append([]int{num}, iiRev...)
	}

	var res float64

	for powOf58, numOfPows := range iiRev {
		y := float64(powOf58)
		res += float64(numOfPows) * math.Pow(58.0, y)
	}

	return int(res)
}
