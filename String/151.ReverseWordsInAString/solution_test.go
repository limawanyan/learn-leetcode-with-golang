package _51_ReverseWordsInAString

import (
	"fmt"
	"testing"
)

func TestReverseWordsInAString(t *testing.T) {
	s := "  Bob    Loves  Alice   "
	fmt.Println(reverseWords(s))
}