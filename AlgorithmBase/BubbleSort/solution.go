package BubbleSort

// 冒泡排序
func BubbleSort(list []int) {
	// 数组长度
	n := len(list)
	// 一轮中是否有元素交换
	didSwap := false
	// 进行n-1轮迭代
	for i := n-1;i > 0;i-- {
		// 每轮都从头开始进行比较，比较到第i位时结束,i位后的元素经过前面的比较已经有序
		for j := 0;j < i;j++ {
			// 如果当前元素比后一个元素大，则交换两元素位置
			if list[j] > list[j + 1] {
				list[j],list[j+1] = list[j+1],list[j]
				didSwap = true
			}
		}
		// 如果没有元素进行交换则说明序列有序，直接返回
		if !didSwap {
			return
		}
	}
}

// 写法2
func BubbleSort2(list []int) {
	// 数组元素个数
	n := len(list)
	// 元素是否已经有序
	flag := false
	// 循环n-1次
	for i := 0;i < n-1;i++{
		flag = false
		for j := 0;j < n-1-i;j++{
			if list[j] > list[j+1]{
				list[j],list[j+1] = list[j+1],list[j]
				flag = true
			}
		}
		// 如果一轮循环中咩有发生交换说明 已经有序
		if !flag{
			break
		}
	}
}