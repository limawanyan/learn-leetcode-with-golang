package _3_RomanToInteger

var romanValues = map[byte]int{'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000}

// 模拟
func romanToInt(s string) (res int) {
	// 字符串长度
	n := len(s)
	// 遍历字符串
	for i := range s{
		// 当前罗马数字对应值
		value := romanValues[s[i]]
		// 当前罗马数字不是最后一个且小于后一个罗马数 -=，否则 +=
		if i < n-1 && value < romanValues[s[i+1]]{
			res -= value
		}else{
			res += value
		}
	}
	return
}

// 写法二 从后往前遍历
func romanToInt2(s string) (res int) {
	// 最后一位罗马数对应数值
	res = romanValues[s[len(s)-1]]
	// 从后往前遍历
	for i := len(s)-1;i > 0;i--{
		// 当前遍历罗马的前一位罗马数对应值
		perNum := romanValues[s[i-1]]
		// 如果前一位罗马小于当前罗马数值 -= 否则 +=
		if perNum < romanValues[s[i]]{
			res -= perNum
		}else{
			res += perNum
		}
	}
	return
}
