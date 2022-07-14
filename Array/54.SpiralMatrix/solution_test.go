package _4_SpiralMatrix

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	list := [][]int{{2,5},{8,4},{0,-1}}
	fmt.Println(spiralOrder4(list))
}
