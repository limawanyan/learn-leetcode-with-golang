package MergeSort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	list := []int{2,46,6,3,23,56,34}
	fmt.Println(MergeSort(list))
}