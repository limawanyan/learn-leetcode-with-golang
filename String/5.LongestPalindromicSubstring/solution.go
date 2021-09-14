package __LongestPalindromicSubstring

// 动态规划
func longestPalindrome(s string) string {
	dp := make([][]bool,len(s))
	// 初始化dp
	for i := 0;i < len(s);i++{
		dp[i] = make([]bool,len(s))
	}
	ans := ""
	for i := 0;i < len(s);i++{
		for k := 0;k <= i;k++{
			dp[i][k] = s[i] == s[k] && (i-1 < k+1 || dp[i-1][k+1])
			if dp[i][k] && i-k+1 > len(ans) {
				ans = s[k:i+1]
			}
		}
	}
	return ans
}

// 中心扩散法 时间复杂度O（n^2）空间复杂度O（1）
//1、围绕中心去扩散，给个字符串加上字符串左右边界，找出最大回文字符串边界值是核心
//2、遍历字符串，处理好奇数和偶数的情况。奇数和偶数的边界值不太一样
//3、找到右-左最大的范围就是最长回文的边界
func longestPalindrome2(s string) string {
	start,end := 0,0
	for i := range s {
		l,r := expand(s,i,i)
		if r - l > end - start{
			start = l
			end = r
		}
		l,r = expand(s,i,i+1)
		if r - l > end - start{
			start = l
			end = r
		}
	}
	return s[start:end+1]
}

func expand(s string,left,right int) (int,int) {
	// 从中心扩散，找到最大回文子串的左右边界
	for ;left >= 0 && right < len(s) && s[left] == s[right];left,right = left-1,right+1{}
	// 去除最后一次循环多余的加减操作
	return left+1,right-1
}