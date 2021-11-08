package main

import (
	"fmt"
)

// 斐波拉契
func main()  {
	n := 0
	fmt.Scan(&n)
	if n < 3 {
		fmt.Println(n)
		return
	}
	var a,b,res int64
	a,b = 1,2
	for i := 3;i <= n;i++ {
		res = (a + b) % int64(1<<29)
		a,b = b,res
	}
	fmt.Println(res)
}