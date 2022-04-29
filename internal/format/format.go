package format

import "fmt"

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
