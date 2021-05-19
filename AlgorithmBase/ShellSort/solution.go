package ShellSort

// 增量序列折半的希尔排序
func ShellSort(list []int)  {
	// 数组长度
	n := len(list)
	// 每次减半，直到步长为1
	for step := n / 2;step >= 1;step /= 2{
		// 开始插入排序，每一轮的步长为step
		for i :=step;i <n;i+=step{
			for j:= i - step;j >= 0;j -= step{
				// 满足插入条件交换元素
				if list[j+step] < list[j] {
					list[j],list[j + step] = list[j + step],list[j]
					continue
				}
				break
			}
		}
	}
}