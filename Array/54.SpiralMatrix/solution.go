package _4_SpiralMatrix

// 模拟 时间 空间复杂度 O（mn）
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 ||  len(matrix[0]) == 0{
		return []int{}
	}
	// 矩阵 行 列
	rows,cols := len(matrix),len(matrix[0])
	// 保存访问路径
	visisted := make([][]bool,rows)
	for i := 0;i < rows;i++{
		visisted[i] = make([]bool,cols)
	}

	var (
		// 总元素个数
		total = rows * cols
		// 保存顺时针旋转结果
		order = make([]int,total)
		// 模拟移动坐标
		row,col = 0,0
		// 移动规则 向右移动 向下移动 向左移动 向上移动
		directions = [][]int{[]int{0,1},[]int{1,0},[]int{0,-1},[]int{-1,0}}
		// 控制移动方向
		directionIndex = 0
	)

	// 模拟移动步数等于矩阵大小
	for i := 0;i < total;i++{
		// 记录当前元素 （0,0）开始
		order[i] = matrix[row][col]
		// 标记已经走过的元素
		visisted[row][col] = true
		// 下一个要走到的坐标
		nextRow,nextCol := row + directions[directionIndex][0],col + directions[directionIndex][1]
		// 如果下一个要走的元素索引越界 或 已经走过了 改变行进方向
		if nextRow < 0  || nextRow >= rows || nextCol < 0 || nextCol >=cols || visisted[nextRow][nextCol]{
			// directionIndex 取值范围 0,1,2,3
			directionIndex = (directionIndex + 1) % 4
		}
		// 下一个元素坐标
		row += directions[directionIndex][0]
		col += directions[directionIndex][1]
	}
	return order
}

// 按层模拟 时间复杂度O（mn）空间复杂度除输出数组O（1）
func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0{
		return []int{}
	}
	var (
		rows,cols = len(matrix),len(matrix[0])
		order = make([]int,rows * cols)
		index = 0
		left,right,top,bottom = 0,cols-1,0,rows-1
	)

	for left <= right && top <= bottom{
		for col := left;col <= right;col++{
			order[index] = matrix[top][col]
			index++
		}
		for row := top + 1;row <= bottom;row++{
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom{
			for col := right - 1;col > left;col--{
				order[index] = matrix[bottom][col]
				index++
			}
			for row := bottom;row > top;row--{
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}

// 方法三
func spiralOrder3(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0{
		return result
	}
	// 左右上下边界
	left,right,top,bottom := 0,len(matrix[0])-1,0,len(matrix)-1

	// 待遍历坐标
	var x,y int
	// 当越界时结束
	for left <= right && top <= bottom{
		// 向右遍历
		for y = left;y <= right && avoid(left,right,top,bottom);y++{
			result = append(result,matrix[x][y])
		}
		// 去除上面多走的一步，回到right
		y--
		// 向下一层
		top++
		// 向下遍历
		for x = top;x <= bottom && avoid(left,right,top,bottom);x++{
			result = append(result,matrix[x][y])
		}
		// 去除多走的一步，回到下轮bottom
		x--
		// 向左移动一格
		right--
		// 向左遍历
		for y = right;y >= left && avoid(left,right,top,bottom);y--{
			result = append(result,matrix[x][y])
		}
		// 去除多走的一步
		y++
		// 向上一层
		bottom--
		// 向上遍历
		for x = bottom;x >= top && avoid(left,right,top,bottom);x--{
			result = append(result,matrix[x][y])
		}
		// 去除多走的
		x++
		// 向右走一格
		left++
	}
	return result
}

func avoid(left, right, top, bottom int) bool {
	return top <= bottom && left <= right
}