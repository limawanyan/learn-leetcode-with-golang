package _6_MinimumWindowSubstring

func minWindow(s string, t string) string {
	// tm存储t串中元素个数，sm存储滑动窗口内元素个数
	tm, sm := map[byte]int{}, map[byte]int{}
	// 匹配的元素个数
	var matchLen int
	// 返回结果
	var ans string

	// 初始化存储t串元素个数
	for i := range t {
		tm[t[i]]++
	}

	// 遍历s串 窗口向右移
	for l, r := 0, 0; r < len(s); r++ {
		// 将s串元素个数存储到sm中
		sm[s[r]]++
		// 如果滑动窗口中s串元素与t元素个数匹配，匹配元素++
		if _, ok := tm[s[r]]; ok && sm[s[r]] == tm[s[r]] {
			matchLen++
		}
		// 成功匹配到最小子串
		for matchLen == len(tm) {
			// 如果是第一次或者本次匹配子串长度小于之前匹配长度则更新
			if len(ans) == 0 || len(s[l:r+1]) < len(ans) {
				ans = s[l : r+1]
			}
			// 如果窗口左边元素时t串元素，且他们匹配的数量正好是相等的
			if _, ok := tm[s[l]]; ok && sm[s[l]] == tm[s[l]] {
				// 匹配元素数量--
				matchLen--
			}
			// 窗口左元素数量--
			sm[s[l]]--
			// 窗口左边向内收缩
			l++
		}
	}
	return ans
}
