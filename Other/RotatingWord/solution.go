package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	str1, _ := reader.ReadString('\n')
	str1 = str1[:len(str1)-1]

	str2, _ := reader.ReadString('\n')
	str2 = str2[:len(str2)-1]

	Check(str1, str2)
}

func Check(str1, str2 string) {
	n1 := len(str1)
	n2 := len(str2)
	// 长度不相等 false
	if n1 != n2 {
		fmt.Println("NO")
		return
	}
	// abcd
	// cdab
	for i := 0; i < n1; i++ {
		for j := i; j <= i+n1; j++ {
			if j == i+n1 {
				fmt.Println("YES")
				return
			} else if str1[j%n1] != str2[j-i] {
				break
			}
		}
	}
	fmt.Println("NO")
}