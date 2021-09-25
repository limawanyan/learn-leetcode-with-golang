package _51_ReverseWordsInAString

func reverseWords(s string) string {
	// 使用双指针删除冗余空格
	slowIndex,fastIndex := 0,0
	// 字符串转byte数组
	b := []byte(s)
	// 删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' '{
		fastIndex++
	}
	// 删除单词间冗余空格
	for ;fastIndex < len(b);fastIndex++{
		// fastIndex-1 > 0 排除第一个元素
		// b[fastIndex - 1] == b[fastIndex] && b[fastIndex] == ' ' 说明有前一个元素和当前元素都是空格
		// 这种情况下 慢指针在第一个空格元素后一个位置不动，快指针继续向右移动
		if fastIndex - 1 > 0 && b[fastIndex - 1] == b[fastIndex] && b[fastIndex] == ' '{
			continue
		}
		// 将fastIndex位置字符替换到slowIndex指针位置
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	// 删除尾部冗余空格
	if slowIndex - 1 > 0 && b[slowIndex - 1] == ' '{
		b = b[:slowIndex - 1]
	}else{
		b = b[:slowIndex]
	}
	// 反转整个字符串
	reverse(&b,0,len(b)-1)
	// 反转单词
	i := 0
	for i < len(b){
		j := i
		// i 为单词起始位置，j 为单词末尾
		for ;j < len(b) && b[j] != ' ';j++{}
		reverse(&b,i,j - 1)
		i = j
		// 跳过空格到达下一个单词开头
		i++
	}
	return string(b)
}

func reverse(b *[]byte,left,right int)  {
	for left < right{
		(*b)[left],(*b)[right] = (*b)[right],(*b)[left]
		left++
		right--
	}
}