package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	start := time.Now()
	data, err := ioutil.ReadFile("primes_0B.bin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count: ", Scan(data))
	fmt.Println(time.Since(start))
}

func Scan(data []byte) (count int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < 8; j++ {
			if GetBit(data[i], j) {
				count++
			}
		}
	}
	return count
}

func Unpack(data []byte) []bool {
	res := make([]bool, len(data)*8)
	for i := 0; i < len(data); i++ {
		for j := 0; j < 8; j++ {
			res[(i*8)+j] = GetBit(data[i], j)
		}
	}
	return res
}

func Between(a, b int) []int {
	return []int{-11, -22, -33}
}

func Is(n int) bool {
	return false
}

func Upto(n int) []int {
	return []int{-11, -22, -33}
}

func GetBit(data byte, index int) bool {
	idx := 7 - index
	bit := data >> idx
	return bit&1 == 1
}
