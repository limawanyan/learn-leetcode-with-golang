package LCP06_GetCoin

import (
	"fmt"
	"testing"
)

func TestMinCoin(t *testing.T) {
	coins := []int{1,5,5,6}
	fmt.Println(minCount(coins))
}