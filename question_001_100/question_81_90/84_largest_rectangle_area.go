package question_81_90

// 84. 柱状图中最大的矩形
// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/

func largestRectangleArea(heights []int) int {
	//return largestRectangleAreaViolent(heights)
	//return largestRectangleAreaDivide(heights)
	return largestRectangleAreaStack(heights)
}

// 1. 暴力破解 依次比较任意两个柱子之间的面积，取最大值
func largestRectangleAreaViolent(heights []int) int {
	var max, min, tmp int
	for i := 0; i < len(heights); i++ {
		min = heights[i]
		for j := i; j < len(heights); j++ {
			if heights[j] < min {
				min = heights[j]
			}
			tmp = min * (j - i + 1)
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}

// 2. 分治 找到最小值，计算最小值左右两边的最大面积以及当前最大面积，取最大值
func largestRectangleAreaDivide(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	}
	var min, index = heights[0], 0
	for i, n := range heights {
		if n < min {
			min, index = n, i
		}
	}
	max := l * min
	if tmp := largestRectangleAreaDivide(heights[:index]); tmp > max {
		max = tmp
	}
	if tmp := largestRectangleAreaDivide(heights[index+1:]); tmp > max {
		max = tmp
	}
	return max
}

// 3. 栈 当需要入栈柱子高度小于栈顶柱子高度时，弹出栈顶柱子，直到栈顶柱子高度小于或等于需要入栈柱子高度
// 弹出栈顶柱子时需计算面积，高度计算方式为当前需要入栈柱子的索引减去栈顶第二个柱子加1索引乘以需要弹出柱子的高度
func largestRectangleAreaStack(heights []int) int {
	heights = append(heights, -1) // 添加-1便于弹出所有柱子
	var stack, l, max = make([]int, len(heights)+1), 0, 0
	stack = append(stack, -1)
	stack[0] = -1
	for i, n := range heights {
		if i != 0 && heights[stack[l]] > n {
			for stack[l] != -1 && n < heights[stack[l]] {
				if (i-stack[l-1]-1)*heights[stack[l]] > max {
					max = (i - stack[l-1] - 1) * heights[stack[l]]
				}
				l--
			}
		}
		l++
		stack[l] = i
	}
	return max
}
