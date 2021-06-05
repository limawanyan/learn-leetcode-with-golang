package _39_SlidingWindowMaximum

import (
	"container/heap"
	"sort"
)

// 暴力破解法(超时时间限制)
func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int,0,k)
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	for i := 0;i <= length-k;i++{
		max := nums[i]
		for j := 1;j < k;j++{
			if max < nums[i + j]{
				max = nums[i+j]
			}
		}
		result = append(result,max)
	}
	return result
}

// 优先队列 大顶堆
var a []int
type hp struct{sort.IntSlice}
func (h *hp) Less(i,j int) bool{
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}
func (h *hp) Push(v interface{})  {
	h.IntSlice = append(h.IntSlice,v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow2(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int,k)}
	for i := 0;i < k;i++{
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int,1,n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k;i < n;i++{
		heap.Push(q,i)
		for q.IntSlice[0] <= i-k{
			heap.Pop(q)
		}
		ans = append(ans,nums[q.IntSlice[0]])
	}
	return ans
}

// 解法三 双端队列
func maxSlidingWindow3(nums []int,k int) []int  {
	if len(nums) == 0 || len(nums) < k{
		return make([]int,0)
	}
	// 存储索引 （最大元素永远保持在左边）
	window := make([]int,0,k)
	result := make([]int,0,len(nums)-k+1)
	for i,v := range nums {
		// 维护索引队列 当前索引大于等于K and 如果元素索引超出了左界
		if i >= k && window[0] <= i-k{
			window = window[1:len(window)]
		}
		// 维护窗口 索引队列不为空 and 当前元素大于索引队列中右边元素
		for len(window) > 0 && nums[window[len(window)-1]] < v{
			// 将右边元素移除掉
			window = window[0:len(window)-1]
		}
		// 存储索引到数组
		window = append(window,i)
		// 当前索引大于等于滑动窗口维护元素个数
		if i >= k-1{
			result = append(result,nums[window[0]])
		}
	}
	return result
}

// 解法四 分块 + 预处理
func maxSlidingWindow4(nums []int, k int) []int {
	n := len(nums)
	prefixMax := make([]int,n)
	suffixMax := make([]int,n)
	for i,v := range  nums{
		if i % k == 0{
			prefixMax[i] = v
		}else {
			prefixMax[i] = max(prefixMax[i-1],v)
		}
	}
	for i := n - 1;i >= 0;i--{
		if i == n-1 || (i+1)%k == 0{
			suffixMax[i] = nums[i]
		}else{
			suffixMax[i] = max(suffixMax[i+1],nums[i])
		}
	}

	ans := make([]int,n-k+1)
	for i := range ans{
		ans[i] = max(suffixMax[i],prefixMax[i+k-1])
	}
	return ans

}

func max(a,b int) int {
	if a > b{
		return a
	}
	return b
}