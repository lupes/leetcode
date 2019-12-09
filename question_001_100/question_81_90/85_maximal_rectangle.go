package question_81_90

// 85. 最大矩形
// https://leetcode-cn.com/problems/maximal-rectangle/
// Topics: 栈 数组 哈希表 动态规划

func maximalRectangle(matrix [][]byte) int {
	height := len(matrix)
	if len(matrix) == 0 {
		return 0
	}
	width := len(matrix[0])
	dp := make([][]int, height+1)
	for i := 0; i <= height; i++ {
		dp[i] = make([]int, width)
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if matrix[j][i] == '1' {
				dp[j+1][i] += dp[j][i] + 1
			}
		}
	}
	max := 0
	for _, i := range dp[1:] {
		tmp := maximalRectangleHelper(i)
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func maximalRectangleHelper(heights []int) int {
	heights = append(heights, -1) // 添加-1便于弹出所有柱子
	var stack, l, max = make([]int, len(heights)+1), 0, 0
	stack = append(stack, -1)
	stack[0] = -1
	for i, n := range heights {
		if i != 0 && heights[stack[l]] > n {
			for stack[l] != -1 && n < heights[stack[l]] {
				tmp := (i - stack[l-1] - 1) * heights[stack[l]]
				if tmp > max {
					max = tmp
				}
				l--
			}
		}
		l++
		stack[l] = i
	}
	return max
}
