package InsertSort

func InsertSort(list []int)  {
	n := len(list)
	// 循环迭代n-1次
	for i := 1;i < n;i++ {
		// 待排序的元素
		deal := list[i]
		// 待排序左边元素
		j := i - 1
		// 如果待排序的元素小于左边已排序元素的最大元素，则进入处理
		if deal < list[j]{
			// 一直往左边查找比较，比排序元素大的元素向右移动
			for ;j >= 0 && deal < list[j];j-- {
				list[j+1] = list[j]
			}
			// 将排序元素插入
			list[j + 1] = deal
		}
	}
}
