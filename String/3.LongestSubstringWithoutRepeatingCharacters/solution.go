package __LongestSubstringWithoutRepeatingCharacters

import (
	"strings"
)

// 只增不减滑动窗口
func lengthOfLongestSubstring(s string) int {
	start,end := 0,0
	for i := 0;i < len(s);i++ {
		// 判断当前循环元素在滑动窗口中是否有重复元素
		index := strings.Index(s[start:i],string(s[i]))
		if index == -1{
			if i+1 > end {
				end = i + 1
			}
		}else{
			// 双指针均指向重复元素后一位
			start += index + 1
			end += index +1
		}
	}
	return end - start
}

// 写法二
func lengthOfLongestSubstring2(s string) int {
	l,r := 0,0
	var result int
	for i := range s{
		index := strings.Index(s[l:i],string(s[i]))
		if index == -1{
			r++
		}else{
			l += index + 1
			r = i + 1
		}
		result = max(len(s[l:r]),result)
	}
	return result
}

// 练习3
func test3(s string) int {
	// 定义滑动窗口
	l,r := 0,0
	// 定义最长子字符串长度
	result := 0
	// 遍历字符串
	for i := range s{
		index := strings.Index(s[l:i],string(s[i]))
		if index == -1{
			r++
		}else {
			l += index+1
			r = i+1
		}
		result = max(result,len(s[l:r]))
	}
	return result
}

func max(x,y int) int {
	if x > y{
		return x
	}else{
		return y
	}
}