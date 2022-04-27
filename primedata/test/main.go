package main

import (
	"fmt"

	"github.com/kevincjackson/smallprimes-go/primedata"
)

func main() {
	fmt.Println(primedata.Is(29))
	fmt.Println(primedata.Upto(100))
	fmt.Println(primedata.Between(50, 100))
}
