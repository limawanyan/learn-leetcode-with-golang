package _0_ValidParentheses

import (
	"strings"
)

// 栈
func isValid(s string) bool {
	// 如果字符串长度为奇数直接返回
	if len(s)%2 != 0 {
		return false
	}
	// 存储左右括号用于后面判断
	pairs := map[byte]byte{
		')':'(',
		'}':'{',
		']':'[',
	}
	// 定义stack用于存储左括号
	var stack []byte
	for i := 0;i < len(s);i++ {
		// 如果字符为右括号
		if pairs[s[i]] > 0{
			// 如果栈为空，或栈顶元素不是与之相匹配的左括号
			if len(stack) == 0 || stack[len(stack) - 1] != pairs[s[i]] {
				return false
			}
			// 将栈顶匹配元素剔除
			stack = stack[:len(stack) - 1]
		}else{
			// 将左括号放入栈
			stack = append(stack, s[i])
		}
	}
	// 判断栈是否为空，如果所有括号匹配则为空
	return len(stack) == 0
}

//替换
func isValid2(s string) bool {
	var length int
	for {
		length = len(s)
		r := strings.NewReplacer("()","","[]","","{}","")
		s = r.Replace(s)
		if length == len(s){
			break
		}
	}
	return len(s) == 0
}