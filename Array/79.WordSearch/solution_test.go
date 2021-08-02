package _9_WordSearch

import (
	"fmt"
	"testing"
)

func TestWordSearch(t *testing.T) {
	list := [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
	fmt.Println(exist(list,"ABCCE"))
}