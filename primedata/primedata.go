package primedata

import (
	_ "embed"
)

//go:embed primes_0B.bin
var data []byte
var MaxInt int = 1_000_000_000

func Between(a, b int) []int {
	primes := []int{}
	for i := a; i <= b; i++ {
		if Is(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func Is(n int) bool {
	if n < 100 {
		return isFromSlice(n)
	} else if n <= MaxInt {
		return isFromRepo(n)
	} else {
		panic("primedata: Input too large. Check the primedata.MaxInt value.")
	}
}

func isFromSlice(n int) bool {
	switch n {
	case 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97:
		return true
	default:
		return false
	}
}

func isFromRepo(n int) bool {
	if n%2 == 0 || n%5 == 0 {
		return false
	} else {
		bucket := (n / 10) * 4
		offset := map[int]int{1: 0, 3: 1, 7: 2, 9: 3}[n%10]
		data1379index := bucket + offset
		byteLocal, bit := data1379index/8, data1379index%8
		return getBit(data[byteLocal], bit)
	}
}

func Upto(n int) []int {
	primes := []int{}
	for i := 2; i <= n; i++ {
		if Is(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func getBit(data byte, index int) bool {
	idx := 7 - index
	bit := data >> idx
	return bit&1 == 1
}
