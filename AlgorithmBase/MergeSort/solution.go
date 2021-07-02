package MergeSort

// 归并排序
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1{
		return nums
	}
	// 分治法 将数组分为两段
	mid := len(nums) / 2
	// 递归处理左右两边
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	// 合并处理结果
	result := merge(left,right)
	return result
}

func merge(left, right []int) (result []int) {
	// 两边数组合并游标
	l := 0
	r := 0
	// 控制索引，只要其中一边元素处理完就停止
	for l < len(left) && r < len(right) {
		// 如果左边元素小于右边元素，则合并到结果集
		if left[l] < right[r]{
			result = append(result,left[l])
			l++
		}else {
			result = append(result,right[r])
			r++
		}
	}
	// 合并剩余部分
	result = append(result,left[l:]...)
	result = append(result,right[r:]...)
	return
}