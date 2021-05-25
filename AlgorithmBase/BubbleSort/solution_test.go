package BubbleSort

import (
	"fmt"
	"testing"
)

func TestBubleSort(t *testing.T) {
	list := []int{1,4,5,6,14,30}
	BubbleSort(list)
	fmt.Print(list)
}