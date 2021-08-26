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

// 写法二
func ShellSort2(list []int)  {
	// 数组长度
	n := len(list)
	// 进行分组,每次对半分组，直到步长为1
	for step := n/2;step >= 1;step /= 2 {
		// 对每一组进行插入排序
		for i := step;i < n;i+=step {
			deal := list[i]
			j := i - step
			for ;j >= 0 && deal < list[j];j -= step {
				list[j + step] = list[j]
			}
			list[j + step] = deal
		}
	}
}

// 练习
func ShellSort3(list []int)  {
	n := len(list)
	for step := n/2;step >= 1;step /= 2{
		for i := step;i < n;i+=step{
			deal := list[i]
			j := i - step
			for ;j>=0 && deal < list[j];j-=step{
				list[j+step] = list[j]
			}
			list[j+step] = deal
		}
	}
}