package __PalindromeNumber

import "strconv"

// 字符串 + 双指针
func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}
	// 转为字符串
	s := strconv.Itoa(x)
	// 左右指针
	l,r := 0,len(s)-1
	for l != r && l < r{
		// 如果出现不相等 false
		if s[l] != s[r]{
			return false
		}
		l++
		r--
	}
	return true
}

// 反转一半数字 时间复杂度：O(logn) 空间复杂度O(1)
func isPalindrome2(x int) bool {
	// x为负数时不是回文数
	// x结尾为0时，只有当x自身为0才算回文数
	if x < 0 || (x%10 == 0 && x != 0){
		return false
	}
	revNumber := 0
	for x > revNumber{
		revNumber = revNumber * 10 + x % 10
		x /= 10
	}
	// 数字长度为奇数时 revNumber去除中位数
	return x == revNumber || x == revNumber / 10
}

