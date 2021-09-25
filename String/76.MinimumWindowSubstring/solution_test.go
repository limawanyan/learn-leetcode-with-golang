package _6_MinimumWindowSubstring

import (
	"fmt"
	"testing"
)

func TestMinimumWindowSubtring(t *testing.T) {
	s := "ADOBECODEBANC"
	s2 := "ABC"
	fmt.Println(minWindow(s,s2))
}