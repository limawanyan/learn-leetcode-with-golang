package _8_ImplementStrstr

// 暴力破解写法一 时间复杂度O(n * m) 空间复杂度O(1)
func strStr(haystack string,needle string) int {
	n,m := len(haystack),len(needle)
	outer:
		for i := 0;i + m <= n;i++ {
			for j := range needle{
				if haystack[i + j] != needle[j]{
					continue outer
				}
			}
			return i
		}
		return -1
}

// 暴力破解写法二
func strStr2(haystack string,needle string)	int {
	if len(needle) == 0{
		return 0
	}
	var j int
	for i := 0;i < len(haystack) - len(needle) + 1;i++{
		for j = 0;j < len(needle);j++ {
			if haystack[i + j] != needle[j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}
	return -1
}

// KMP算法 时间复杂度 O(m + n) 空间复杂度 O(m) m是字符串 needle 的长度
func strStr3(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}