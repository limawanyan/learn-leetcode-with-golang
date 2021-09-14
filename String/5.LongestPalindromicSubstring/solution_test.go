package __LongestPalindromicSubstring

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "abcd"
	for i := range s{
		fmt.Println(string(s[i]))
	}
}