package _6_Permutations

func permute(nums []int) [][]int {
	result := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int){
		// 结束条件
		if len(path) == len(nums){
			result = append(result,path)
			return
		}
		for _,value := range nums {
			// 如果存在这个数以及这个数已经访问过了跳过
			if setStaus,exists := visited[value];exists && setStaus{
				continue
			}
			// 设置为已访问状态
			visited[value] = true
			// 递归添加元素
			dfs(append(path,value))
			// 回溯
			visited[value] = false
		}
	}
	dfs([]int{})
	return result
}