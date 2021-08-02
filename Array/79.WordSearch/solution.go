package _9_WordSearch

// dfs + 辅助数组
func exist(board [][]byte, word string) bool {
	m,n := len(board),len(board[0])
	// 访问记录
	visited := make([][]bool,m)
	for i := 0;i < m;i++{
		visited[i] = make([]bool,n)
	}

	var canFind func(r,c,i int) bool
	canFind = func(r, c, i int) bool {
		// 已经找到复合字符串长度的路径
		if i == len(word){
			return true
		}
		// 索引越界
		if r < 0 || r >= m || c < 0 || c >= n{
			return false
		}
		// 当前元素已访问过或者当前元素不等于当前需要找的字符
		if visited[r][c] || board[r][c] != word[i] {
			return false
		}
		// 设置当前元素为已访问状态
		visited[r][c] =  true
		// 递归剩余路径 在 || 判断中，只要有符合的便不会继续执行后面的判断，达到剪枝的目的
		if canFind(r+1,c,i+1) || canFind(r-1,c,i+1) || canFind(r,c+1,i+1) || canFind(r,c-1,i+1){
			return true
		}
		// 没有符合的路径 将当前元素访问状态回退
		visited[r][c] = false
		return false
	}

	for i := 0;i < m;i++ {
		for j := 0;j < n;j++{
			if board[i][j] == word[0] && canFind(i,j,0){
				return true
			}
		}
	}
	return false
}

// dfs + 直接修改原数组
func exist2(board [][]byte, word string) bool {
	m,n := len(board),len(board[0])
	var canFind func(r,c,i int) bool
	canFind = func(r, c, i int) bool {
		// 已经找到复合字符串长度的路径
		if i == len(word){
			return true
		}
		// 索引越界
		if r < 0 || r >= m || c < 0 || c >= n{
			return false
		}
		// 如果当前元素不等于寻找元素
		if board[r][c] != word[i] {
			return false
		}
		// 修改元素，表示已访问
		temp := board[r][c]
		board[r][c] =  ' '
		// 递归剩余路径 在 || 判断中，只要有符合的便不会继续执行后面的判断，达到剪枝的目的
		if canFind(r+1,c,i+1) || canFind(r-1,c,i+1) || canFind(r,c+1,i+1) || canFind(r,c-1,i+1){
			return true
		}
		// 没有符合的路径 将当前元素回退
		board[r][c] =  temp
		return false
	}

	for i := 0;i < m;i++ {
		for j := 0;j < n;j++{
			if board[i][j] == word[0] && canFind(i,j,0){
				return true
			}
		}
	}
	return false
}