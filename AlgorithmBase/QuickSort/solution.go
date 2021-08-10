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

// 练习2 将数组进行快速排序，返回排序后的中间元素索引
func partition2(list []int,begin int,end int) int {
	// list[being]为基准元素，用来比较
	// 定义左右指针
	left := begin + 1
	right := end
	// 当左右两指针重合时结束
	for left < right{
		// 判断左指针元素是否大于基准元素，如果大于基准元素则左右指针元素交换位置，右指针--，否则左指针++
		if list[left] > list[begin] {
			list[left],list[right] = list[right],list[left]
			right--
		}else{
			left++
		}
	}

	// 当前左右指针重合，指向元素左边的都是小于基准元素的，右边的都是大于基准元素的
	// 这里需要判断，指向元素是否比基准元素大，如果大于等于基准元素，那么left左移一格，将基准元素与比他小的元素替换位置
	// 如果小于基准元素，那么不需要移动，直接将指向元素和基准元素替换位置即可
	if list[left] >= list[begin]{
		left--
	}

	list[begin],list[left] = list[left],list[begin]

	return left
}

// 快速排序练习
func QuickSort2(list []int,begin int,end int)  {
	if begin < end{
		// 切分排序
		loc := partition2(list,begin,end)
		// 递归排序左边
		QuickSort2(list,begin,loc-1)
		// 递归排序右边
		QuickSort2(list,loc+1,end)
	}
}

func QuickSort3(list []int,begin,end int){
	if begin < end{
		loc := partition3(list,begin,end)
		QuickSort3(list,begin,loc-1)
		QuickSort3(list,loc+1,end)
	}
}

func partition3(list []int,begin,end int) int {
	left := begin+1
	right := end
	for left < right{
		if list[left] > list[begin]{
			list[left],list[right] = list[right],list[left]
			right--
		}else {
			left++
		}
	}
	if list[begin] <= list[left]{
		left--
	}
	list[begin],list[left] = list[left],list[begin]
	return left
}