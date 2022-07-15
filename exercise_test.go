package learn_leetcode_with_golang

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	wf := map[string]int{}
	a := "abcdefg"
	for j, n := 1, len(a); j <= n; j++ {
		for k := 0; k < n; k++ {
			wf[a[:j]+"#"+a[k:]] = 0
		}
	}
	fmt.Println(wf)
}
