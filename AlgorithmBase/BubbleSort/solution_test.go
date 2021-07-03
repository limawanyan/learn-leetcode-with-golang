package BubbleSort

import (
	"fmt"
	"testing"
)

func TestBubleSort(t *testing.T) {
	list := []int{8,4,15,6,14,30}
	BubbleSort(list)
	fmt.Print(list)
}