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

// 写法二
func romanToInt2(s string) (res int) {
	res = romanValues[s[len(s)-1]]
	for i := len(s)-1;i > 0;i--{
		perNum := romanValues[s[i-1]]
		if perNum < romanValues[s[i]]{
			res -= perNum
		}else{
			res += perNum
		}
	}
	return
}
