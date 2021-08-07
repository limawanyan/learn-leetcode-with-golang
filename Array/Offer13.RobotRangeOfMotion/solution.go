package Offer13_RobotRangeOfMotion

// 计算位数和
func sums(x int) (res int) {
	for x != 0{
		res += x % 10
		x = x / 10
	}
	return res
}

// DFS
func movingCount(m int, n int, k int) int {
	// 初始化矩阵判断访问状态
	visited := make([][]bool,m)
	for i := 0;i < m;i++{
		visited[i] = make([]bool,n)
	}

	// dfs深度优先遍历
	var dfs func(i,j,si,sj int) int
	dfs = func(i, j, si, sj int) int {
		// 判断索引越界，位数和符合，是否已访问
		if i >= m || j >= n || si + sj > k || visited[i][j]{
		return 0
	}
		visited[i][j] = true
		// 计算右和下边的位数和
		sj1 := sj + 1
		si1 := si + 1
		if (j + 1) % 10 == 0{
		sj1 = sj - 8
	}
		if (i + 1) % 10 == 0{
		si1 = si - 8
	}
		return 1 + dfs(i,j+1,si,sj1) + dfs(i+1,j,si1,sj)
	}

	return dfs(0,0,0,0)
}

// BFS
func movingCount2(m int, n int, k int) int {
	// 队列
	var queue [][4]int
	visited := make([][]bool,m)
	for i := 0;i < m;i++{
		visited[i] = make([]bool,n)
	}
	queue = append(queue,[4]int{0,0,0,0})
	res := 0
	for len(queue) > 0{
		back := queue[0]
		queue = queue[1:]
		i := back[0]
		j := back[1]
		si := back[2]
		sj := back[3]
		if i >= m || j >= n || si + sj > k || visited[i][j]{
			continue
		}
		res++
		visited[i][j] = true

		sj1 := sj + 1
		si1 := si + 1
		if (j+1)%10 == 0{
			sj1 = sj - 8
		}
		if (i+1)%10 == 0 {
			si1 = si - 8
		}
		queue = append(queue,[4]int{i+1,j,si1,sj})
		queue = append(queue,[4]int{i,j+1,si,sj1})
	}
	return res
}