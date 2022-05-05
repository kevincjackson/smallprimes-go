package format

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kevincjackson/smallprimes-go/pkg/primedata"
)

type sliceFormat struct {
	Prefix    string
	Separator string
	Suffix    string
}

var formatMap = map[string]sliceFormat{
	"json":     {"[", ", ", "]"},
	"newlines": {"", "\n", ""},
	"spaces":   {"", " ", ""},
}

func PrintInts(xs []int, style string) {
	s, ok := formatMap[style]
	if !ok {
		s = formatMap["spaces"]
	}
	for i, x := range xs {
		if i == 0 {
			fmt.Printf(s.Prefix)
		}
		fmt.Printf("%d", x)
		if i != len(xs)-1 {
			fmt.Printf(s.Separator)
		} else {
			fmt.Printf(s.Suffix)
		}
	}
	fmt.Println()
}

// 10_000_000 => ~266ms (twice as fast as int -> string -> join)
// Get rid of the last comma?
func Between(x int, y int, prefix string, separator string, suffix string) string {
	builder := strings.Builder{}
	builder.Grow(y - x)
	builder.WriteString(prefix)
	if x < 2 {
		x = 2
	}
	for i := x; i <= y; i++ {
		if primedata.Is(i) {
			builder.WriteString(strconv.Itoa(i))
			if i < y-1 {
				builder.WriteString(separator)
			}
		}
	}
	builder.WriteString(suffix)
	return builder.String()
}
