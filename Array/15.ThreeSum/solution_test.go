package _5_ThreeSum

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	list := []int{ -1, 0, 1, 2, -1, -4}
	result := threeSum2(list)
	fmt.Println(result)
}
