package SimpleSelectionSort

import (
	"fmt"
	"testing"
)

func TestSimpleSelectionSort(t *testing.T) {
	list := []int{45,5,48,40,22,20}
	//SimpleSelectionSort(list)
	//fmt.Println(list)
	//
	//list = []int{45,80,48,40,22,78}
	//SelectGoodSort(list)
	//fmt.Println(list)

	SelectSort(list)
	fmt.Println(list)
}
