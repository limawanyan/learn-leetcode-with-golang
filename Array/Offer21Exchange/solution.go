package Offer21

// 左右指针
func exchange(nums []int) []int {
	if len(nums) <= 1{
		return nums
	}
	// 定义左右两个指针
	left,right := 0,len(nums)-1
	// 两指针相遇是循环结束
	for left < right{
		// 如果左指针元素为偶数，右指针为奇数，两元素进行交换
		if nums[left]%2==0&&nums[right]%2!=0{
			nums[left],nums[right]=nums[right],nums[left]
		}
		// 左指针一直往右移动，直到为奇数
		for left < right && nums[left] % 2 != 0{
			left++
		}
		// 右指针一直往左移动，直到为偶数
		for left < right && nums[right] % 2 == 0{
			right--
		}
	}
	return nums
}

// 快慢指针
func exchange2(nums []int) []int {
	if len(nums) <= 1{
		return nums
	}
	// 定义两个快慢指针，快指针先走，左指针指向待替换的元素，快指针找到下一个需要替换的奇数
	for low,fast := 0,0;fast < len(nums);fast++{
		// 当快指针找到奇数时进行替换
		if nums[fast] % 2 != 0{
			nums[low],nums[fast] = nums[fast],nums[low]
			low++
		}
	}
	return nums
}