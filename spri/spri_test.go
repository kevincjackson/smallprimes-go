package spri

import (
	"testing"
)

func TestIs0(t *testing.T) {
	AssertEqual(t, "Is(0)", Is(0), false)
}

func TestIs2(t *testing.T) {
	AssertEqual(t, "Is(2)", Is(2), true)
}

func TestIs3(t *testing.T) {
	AssertEqual(t, "Is(3)", Is(3), true)
}

func TestIs4(t *testing.T) {
	AssertEqual(t, "Is(4)", Is(4), false)
}

func TestIs101(t *testing.T) {
	AssertEqual(t, "Is(101)", Is(101), true)
}

func TestIs103(t *testing.T) {
	AssertEqual(t, "Is(103)", Is(103), true)
}

func TestIs105(t *testing.T) {
	AssertEqual(t, "Is(105)", Is(105), false)
}

func TestIs107(t *testing.T) {
	AssertEqual(t, "Is(107)", Is(107), true)
}

func TestIs999999937(t *testing.T) {
	AssertEqual(t, "Is(999999937)", Is(999999937), true)
}

func TestIs999999997(t *testing.T) {
	AssertEqual(t, "Is(999999997)", Is(999999997), false)
}

func AssertEqual(t *testing.T, name string, act interface{}, exp interface{}) {
	if act != exp {
		t.Log(name, "expected", exp, "got", act)
		t.Fail()
	}
}
