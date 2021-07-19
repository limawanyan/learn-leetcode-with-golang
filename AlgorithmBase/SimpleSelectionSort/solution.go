package SimpleSelectionSort

// 普通简单选择排序
func SimpleSelectionSort(list []int)  {
	n := len(list)
	// 进行n-1轮迭代
	for i := 0;i < n-1;i++ {
		// 每次从第i位开始找最小元素
		min := list[i]
		minIndex := i
		for j := i + 1;j < n;j++{
			// 找到更小的元素则更改
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}
		if i != minIndex{
			list[i],list[minIndex] = list[minIndex],list[i]
		}
	}
}

// 优化版简单选择排序
func SelectGoodSort(list []int)  {
	n := len(list)
	for i := 0;i < n/2;i++{
		minIndex := i // 最小值下标
		maxIndex := i // 最大值下标

		// 在这一轮迭代中找到最小和最大值下标
		for j := i + 1;j < n-i;j++ {
			if list[j] > list[maxIndex] {
				maxIndex = j
				continue
			}
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}

		// 最大值是开头元素，最小值不是末尾元素，先将最大值与末尾元素替换，再将最小元素与开头元素替换
		if maxIndex == i && minIndex != n-i-1 {
			list[maxIndex],list[n-i-1] = list[n-i-1],list[maxIndex]
			list[i],list[minIndex] = list[minIndex],list[i]
		}else if maxIndex == i && minIndex == n-i-1 {
			// 最大元素在开头，最小元素再末尾，直接交换
			list[minIndex],list[maxIndex] = list[maxIndex],list[minIndex]
		}else {
			// 否则先将最小元素放开头，再将最大值放末尾
			list[i],list[minIndex] = list[minIndex],list[i]
			list[n-i-1],list[maxIndex] = list[maxIndex],list[n-i-1]
		}
	}
}

// 选择排序 时间复杂度
func SelectSort(list []int){
	// 元素个数
	n := len(list)
	// 循环n-1次
	for i := 0;i < n-1;i++ {
		// 最小元素
		min := i
		for j := i+1;j < n;j++ {
			if list[j] < list[min]{
				min = j
			}
		}
		list[i],list[min] = list[min],list[i]
	}
}