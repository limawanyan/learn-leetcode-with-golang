package LetterCombinations

import (
	"fmt"
	"testing"
)

func TestLetter(t *testing.T) {
	var result []string
	LetterCombinations(&result,"","112")
	fmt.Println(result)
}