package cryptoClass

import (
	"./lib"
	"fmt"
)

func main() {
	s := crypto.Encode("张学友")
	fmt.Println(s)
	fmt.Println(crypto.Decode(s))
}
