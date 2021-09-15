package _00_NumberOfIslands

import (
	"fmt"
	"testing"
)

func TestNumberOflslands(t *testing.T) {
	//lands := [][]byte{
	// {'1','1','1','1','0'},
	// {'1','1','0','1','0'},
	// {'1','1','0','0','0'},
	// {'0','0','0','1','0'},
	//}

	lands2 := [][]byte{
		{'1'},
	}
	fmt.Println(numIslands(lands2))

}
