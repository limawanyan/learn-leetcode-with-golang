package QuickSort

// 普通快速排序
func QuickSort(array []int, begin, end int) {
	if begin < end{
		// 进行切分
		loc := partition(array,begin,end)
		// 对左部分进行快排
		QuickSort(array,begin,loc - 1)
		// 对右部分进行快排
		QuickSort(array,loc + 1,end)
	}
}

// 切分函数，并返回切分元素的下标
func partition(array []int,begin,end int) int {
	i := begin + 1	// 将array[begin]作为基数，因此从array[begin+1]开始与基准数比较
	j := end
	// 重合时结束
	for i < j{
		if array[i] > array[begin]{
			array[i],array[j] = array[j],array[i]
			j--
		}else{
			i++
		}
	}

	// 循环后i=j
	if array[i] >= array[begin]{
		i--
	}

	array[begin],array[i] = array[i],array[begin]
	return i
}