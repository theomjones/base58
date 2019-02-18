# base58
Encode and decode integers to and from Flickr flavoured base58, useful for short URLs.
Alphanumeric reference: 123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz

## Install
`go get github.com/theomjones/base58`

## Usage

```go
package main

import (
	"fmt"
	"github.com/theomjones/base58"
)

func main() {
	e := base58.Encode(90218978)
	fmt.Println(e) // "8yPxd"
	d := base58.Decode("8yPxd")
	fmt.Println(d) // 90218978
}
```