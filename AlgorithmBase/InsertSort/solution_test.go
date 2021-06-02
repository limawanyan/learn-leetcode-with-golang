package InsertSort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	list := []int{30,76,34,87,55,63}
	InsertSort(list)
	fmt.Println(list)
}