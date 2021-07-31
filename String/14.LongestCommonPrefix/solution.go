package _4_LongestCommonPrefix

// 横向搜索
func longestCommonPrefix(strs []string) string {
	// 如果字符数组为空直接返回
	if len(strs) == 0{
		return ""
	}
	// 公共前缀,初始化为第一个
	prefix := strs[0]
	// 数组元素个数
	count := len(strs)
	// 循环获取公共前缀,将前一个与后一个元素比较
	for i := 1;i < count;i++{
		prefix = lcp(prefix,strs[i])
		// 如果得到的最新公共前缀长度为0，也说明整个数组内元素没有公共前缀直接返回
		if len(prefix) == 0{
			break
		}
	}
	return prefix
}

// 计算最长公共前缀
func lcp(str1, str2 string) string {
	// 最短元素长度，因为最长公共前缀最长为最短元素长度
	length := min(len(str1),len(str2))
	// 比较起始索引
	index := 0
	// 索引没有越界 且 两字符串索引处元素相等
	for index < length && str1[index] == str2[index]{
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x > y{
		return y
	}
	return x
}

// 纵向搜索
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	// col
	for i := 0;i <len(strs[0]);i++{
		// row
		for j := 1;j < len(strs);j++{
			// 如果列已经到最后一个元素 或者 j行i列与第一行i列元素不同则返回前面相同的前缀
			if i == len(strs[j]) || strs[j][i] != strs[0][i]{
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 二分法
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	// 判断各元素指定长度前缀是否相等
	isCommonPrefix := func(length int) bool {
		str0,count := strs[0][:length],len(strs)
		for i := 1;i < count;i++{
			if strs[i][:length] != str0{
				return false
			}
		}
		return true
	}
	// 遍历数组获取到最小元素长度
	minLength := len(strs[0])
	for _,s := range strs{
		if len(s) < minLength{
			minLength = len(s)
		}
	}
	// 从0开始,最长公共前缀长度为minLength
	low,high := 0,minLength
	for low < high{
		// 计算中间值
		mid := (high - low + 1) / 2 + low
		if isCommonPrefix(mid){
			low = mid
		}else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}

