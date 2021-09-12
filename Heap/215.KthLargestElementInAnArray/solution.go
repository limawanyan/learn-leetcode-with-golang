package _15_KthLargestElementInAnArray

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
)

// 解法一 排序，排序的方法反而速度是最快的
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 解法二 快速选择排序 这个方法的理论依据是 partition 得到的点的下标就是最终排序之后的下标，根据这个下标，我们可以判断第 K 大的数在哪里
// 时间复杂度 O(n)，空间复杂度 O(log n)，最坏时间复杂度为 O(n^2)，空间复杂度 O(n)
func findKthLargest2(nums []int, k int) int {
	m := len(nums) - k + 1 // mth smallest, from 1..len(nums)
	return selectSmallest(nums, 0, len(nums)-1, m)
}

func selectSmallest(nums []int, l, r, i int) int {
	if l >= r {
		return nums[l]
	}
	q := partition(nums, l, r)
	k := q - l + 1
	if k == i {
		return nums[q]
	}
	if i < k {
		return selectSmallest(nums, l, q-1, i)
	} else {
		return selectSmallest(nums, q+1, r, i-k)
	}
}

func partition(nums []int, l, r int) int {
	k := l + rand.Intn(r-l+1) // 此处为优化，使得时间复杂度期望降为 O(n)，最坏时间复杂度为 O(n^2)
	nums[k], nums[r] = nums[r], nums[k]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

// 解法三 优先队列 小顶堆 O(nlogn)
type TopList []int

func (t TopList)Len()int{
	return len(t)
}

func (t TopList)Swap(i,j int){
	t[i],t[j] = t[j],t[i]
}

func (t TopList)Less(i,j int)bool{
	return  t[i]<t[j]
}

func (t *TopList)Push(x interface{}){
	*t = append(*t,x.(int))
}

func (t *TopList)Pop()interface{}{
	x := (*t)[len(*t)-1]
	*t= (*t)[:len(*t)-1]
	return x
}

func findKthLargest3(nums []int, k int) int {
	m :=make(TopList,0)
	size := 0
	for i := range nums{
		if size<k{
			heap.Push(&m,nums[i])
			size++
		}else {
			// 判断堆顶元素是否小于当前元素
			if nums[i]>m[0]{
				heap.Pop(&m)
				heap.Push(&m,nums[i])
			}
		}
	}
	fmt.Println(m)
	return m[0]
}


// 堆排
var kk int
func findKthLargest4(nums []int,k int) int {
	kk = k
	heapSort(nums)
	return nums[len(nums)-k]
}

func heapSort(nums []int)  {
	end := len(nums)-1
	for i := end / 2;i >= 0;i--{
		sink(nums,i,end)
	}
	for i:=end;i>=0;i--{
		nums[i],nums[0]= nums[0],nums[i]
		if kk == len(nums) - end{
			return
		}
		end--
		sink(nums,0,end)
	}
}

func sink(nums[]int,root,end int)  {
	for{
		child := root*2 + 1
		if child > end{
			return
		}
		if child < end && nums[child] <= nums[child+1]{
			child++
		}
		if nums[root] > nums[child]{
			return
		}
		nums[root],nums[child] = nums[child],nums[root]
		root = child
	}
}