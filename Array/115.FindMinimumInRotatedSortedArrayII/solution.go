package _15_FindMinimumInRotatedSortedArrayII

// 暴力法
func minArray(numbers []int) int {
	n := len(numbers)
	for i:=1;i < n;i++{
		if numbers[i] < numbers[i-1]{
			return numbers[i]
		}
	}
	return numbers[0]
}

// 二分法
func minArray2(numbers []int) int {
	left := 0
	right := len(numbers)-1
	for left < right{
		mid := left + (right - left) / 2
		// 如果中间元素大于最右边元素[4、5、6、1、2、3] 最小元素在右边 且不可能为中间元素
		if numbers[mid] > numbers[right]{
			left = mid + 1
			// 如果中间元素小于最右元素[4、5、1、2、3] 最小元素在左边 可能包含中间元素
		}else if numbers[mid] < numbers[right]{
			right = mid
		}else{
			// 相等则抛弃右边元素
			right--
		}
	}
	return numbers[left]
}