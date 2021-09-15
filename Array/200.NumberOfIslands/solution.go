package _00_NumberOfIslands

// DFS 时间复杂度 O（MN） 空间复杂度 O（MN） 最坏情况下，整个网格都是陆地，dfs深度达到MN
func numIslands(grid [][]byte) int {
	res := 0
	row := len(grid)
	col := len(grid[0])
	var dfs func(grid [][]byte,i,j int)
	dfs = func(grid [][]byte,i,j int) {
		if i < 0 || j < 0 || i >= row || j >= col || grid[i][j] == '0'{
			return
		}
		if grid[i][j] == '1'{
			grid[i][j] = '0'
		}
		dfs(grid,i-1,j)
		dfs(grid,i+1,j)
		dfs(grid,i,j-1)
		dfs(grid,i,j+1)
	}
	for i := 0;i < row;i++{
		for j := 0;j < col;j++{
			if grid[i][j] == '1'{
				res++
				dfs(grid,i,j)
			}
		}
	}
	return res
}

// BFS 时间复杂度 O（MN） 空间复杂度 O（min（M,N））
// 最坏情况下，整个网格为陆地，队列的大小可以达到min（M,N）
func numIslands2(grid [][]byte) int {
	res := 0
	row := len(grid)
	col := len(grid[0])
	var bfs func(grid [][]byte,i,j int)
	bfs = func(grid [][]byte, i, j int) {
		queue := [][]int{[]int{i,j}}
		for len(queue) != 0{
			cur := queue[0]
			queue = queue[1:]
			i,j = cur[0],cur[1]
			if i >= 0 && i < row && j >= 0 && j < col && grid[i][j] == '1'{
				grid[i][j] = '0'
				queue = append(queue,[]int{i+1,j})
				queue = append(queue,[]int{i-1,j})
				queue = append(queue,[]int{i,j+1})
				queue = append(queue,[]int{i,j-1})
			}
		}
	}
	for i := 0;i < row;i++ {
		for j := 0;j < col;j++{
			if grid[i][j] == '1'{
				bfs(grid,i,j)
				res++
			}
		}
	}
	return res
}