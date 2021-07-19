package _8_MergeSortedArray

import "sort"

// 方法一：直接合并排序
func merge(nums1 []int,m int,nums2 []int,n int)  {
	copy(nums1[m:],nums2)
	sort.Ints(nums1)
}

// 方法二：双指针 + 辅助数组
func merge2(nums1 []int,m int,nums2 []int,n int)  {
	// 定义两指针分别指向头部
	p1,p2 := 0,0
	// 定义辅助数组
	sorted := make([]int,0,m+n)
	// 当某个数组已经指向最后一个元素时结束循环
	for {
		// 当数组1已经指向最后一个元素，将数组2剩余元素添加进辅助数组
		if p1 == m {
			sorted = append(sorted,nums2[p2:]...)
			break
		}
		// 当数组2已经指向最后一个元素，将数组1剩余元素添加进辅助数组
		if p2 == n {
			sorted = append(sorted,nums1[p1:]...)
			break
		}
		// 将小的元素添加到复制数组，相应指针索引+1
		if nums1[p1] < nums2[p2]{
			sorted = append(sorted,nums1[p1])
			p1++
		}else{
			sorted = append(sorted,nums2[p2])
			p2++
		}
	}
	copy(nums1,sorted)
}


// 方法三：逆向双指针
func merge3(nums1 []int,m int,nums2 []int,n int)  {
	// p 是nums1数组长度 用于逆序存放元素
	for p := m + n;m > 0 && n > 0;p--{
		// 如果nums2末尾元素大于nums1末尾元素，将nums2末尾元素设置到nums1最后，nums2指针--
		if nums1[m-1] <= nums2[n-1]{
			nums1[p-1] = nums2[n-1]
			n--
		}else{
			nums1[p-1] = nums1[m-1]
			m--
		}
	}
	// 如果nums2还有剩余元素，全部逆序放入nums1
	for ;n > 0;n--{
		nums1[n-1] =  nums2[n-1]
	}
}