package _8_ImplementStrstr

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	haystack,needle := "hello","ll"
	fmt.Println(strStr(haystack,needle))
	fmt.Println(strStr2(haystack,needle))
	fmt.Println(strStr3(haystack,needle))
}

func TestKMP(t *testing.T)  {
	needle := "aabaac"
	pi := make([]int,len(needle) )
	// 循环len(needle)次
	for i, j := 1, 0; i < len(needle); i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	fmt.Println(pi)
}