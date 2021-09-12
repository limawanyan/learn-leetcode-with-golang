package _6_Permutations

import (
	"fmt"
	"testing"
)

func TestPernutations(t *testing.T) {
	nums := []int{1,2,3}
	fmt.Println(permute(nums))
}

func TestName(t *testing.T) {
	visited := map[int]bool{}
	visited[1] = true
	visited[0] = true
	a,b := visited[1]
	c,d := visited[3]
	fmt.Println(a,b)
	fmt.Println(c,d)
}