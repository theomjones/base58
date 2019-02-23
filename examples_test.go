package base58_test

import (
	"fmt"

	"github.com/theomjones/base58"
)

func ExampleEncode() {
	s := base58.Encode(100)
	fmt.Println(s)
	// Output: 2j
}

func ExampleDecode() {
	i := base58.Decode("2j")
	fmt.Println(i)
	// Output: 100
}
