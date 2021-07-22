package ShellSort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	list := []int{5}
	ShellSort2(list)
	fmt.Println(list)

	list1 := []int{5,9}
	ShellSort2(list1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	ShellSort2(list2)
	fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 2, 4, 23, 467, 85, 23, 567, 335, 677, 33, 56, 2, 5, 33, 6, 8, 3}
	ShellSort2(list3)
	fmt.Println(list3)
}