package _40_SearchA2DMatrixII

// 暴力破解
func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])
	for i := 0;i < row;i++{
		for j := 0;j < col;j++{
			if matrix[i][j] == target{
				return true
			}
		}
	}
	return false
}

// 搜索 以左下角为起点
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0{
		return false
	}
	row := len(matrix)-1
	col := len(matrix[0])
	index := 0
	// 左下角为起点进行搜索，行列索引越界时结束
	for row >= 0 && index < col{
		// 当前元素等于搜索元素返回true
		if matrix[row][index] == target{
			return true
		}
		// 如果当前元素大于搜索元素，向上移动一行，否则向右移动一列
		if matrix[row][index] > target {
			row--
		}else{
			index++
		}
	}
	return false
}