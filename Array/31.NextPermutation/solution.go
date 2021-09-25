package _1_NextPermutation

func nextPermutation(nums []int)  {
	n := len(nums)
	// 从后往前遍历
	i := n-2
	// 如果当前元素大于或等于后一个元素，说明是降序的
	// 执行 i-- 直到找到升序元素
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// 从后往前查找比i元素小的元素
		j := n-1
		for j >= 0 && nums[i] >= nums[j]{
			j--
		}
		// 交换元素位置
		nums[i],nums[j] = nums[j],nums[i]
	}
	// 将i后面部分元素反转，从降序变为升序
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}