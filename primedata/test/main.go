package main

import (
	"fmt"

	"github.com/kevincjackson/smallprimes-go/primedata"
)

func main() {
	fmt.Println(primedata.Upto(999_000_000))
}
