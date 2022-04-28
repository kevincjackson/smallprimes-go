package main

import (
	"fmt"

	"github.com/kevincjackson/smallprimes-go/primedata"
)

func main() {
	res := primedata.Between(999_000_000, 1_000_000_000)
	fmt.Println(res)
	fmt.Println("count:", len(res))
}
