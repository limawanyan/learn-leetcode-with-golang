package _45_PrefixAndSuffixSearch

import (
	"fmt"
	"testing"
)

func TestSolution1(t *testing.T) {
	words := []string{"apple"}
	obj := Constructor1(words)
	fmt.Println(obj.F1("a", "e"))
}

func TestSolution2(t *testing.T) {
	words := []string{"apple"}
	obj := Constructor2(words)
	fmt.Println(obj.F2("a", "e"))
}
