package __LongestSubstringWithoutRepeatingCharacters

import (
	"strings"
)


// 只增不减滑动窗口 最优
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

// 动态规划
func lengthOfLongestSubstring3(s string) int {
	if len(s)<=1{
		return len(s)
	}
	max:=1
	// dp[i]表示以s[i]结尾的最长不重复串长度
	dp :=make([]int,len(s))
	dp[0]=1    //边界
	for i:=1;i<len(s);i++{
		for j:=i-dp[i-1];j<i;j++{
			if s[j]==s[i]{   //如果有相同的，则重新计算长度
				dp[i]=i-j
				goto LOOP
			}
		}
		dp[i]=dp[i-1]+1   //否则就长度+1
	LOOP:
		if dp[i]>max{
			max=dp[i]
		}
	}
	return max
}

// 滑动窗口 哈希MAP + 双指针 时间复杂度O（n）左右指针遍历字符串一次 空间复杂度O（字符长度）
func lengthOfLongestSubstring4(s string) int {
	// 使用哈希集合记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为-1，相当于在字符串的左边界的左侧，还没开始移动
	rk,ans := -1,0
	// 外层循环控制窗口的左边界相当于左指针
	for i:=0;i<n;i++{
		if i!=0{
			// 左指针向右移动一格，移除一格字符
			delete(m,s[i-1])
		}
		// 不断移动右指针并记录到map,当碰到重复或到最右端结束
		for rk + 1 < n && m[s[rk+1]] == 0{
			m[s[rk+1]]++
			rk++
		}
		// 第i到rk个字符是一个最长的无重复字符子串
		ans = max(ans,rk-i+1)
	}
	return ans
}
