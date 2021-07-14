package _5_ThreeSum

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	list := []int{2,-5,6,0,-2,-5,10}
	result := threeSum(list)
	fmt.Println(result)
}
