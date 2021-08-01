package _40_SearchA2DMatrixII

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	matrix := [][]int{
		[]int{1,4,7,11,15},
		[]int{2,5,8,12,19},
		[]int{3,6,9,16,22},
		[]int{10,13,14,17,24},
		[]int{18,21,23,26,30},
	}
	fmt.Println(searchMatrix2(matrix,5))
}
