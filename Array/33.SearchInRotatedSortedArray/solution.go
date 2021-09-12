package _3_SearchInRotatedSortedArray

// 二分法 时间复杂度 O（logn）每次至少能排除四分之一的元素 空间复杂度O（1）
func search(nums []int, target int) int {
	left,right := 0,len(nums)-1
	for left <= right{
		mid := (right - left) / 2 + left
		if nums[mid] == target{
			return mid
		}
		if nums[mid] >= nums[left]{
			if nums[mid] > target && target >= nums[left]{
				right = mid - 1
			}else{
				left = mid + 1
			}
		}else {
			if nums[mid] < target && target <= nums[right]{
				left = mid + 1
			}else{
				right = mid - 1
			}
		}
	}
	return -1
}

