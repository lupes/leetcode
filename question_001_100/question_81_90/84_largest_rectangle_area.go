package question_81_90

// 84. 柱状图中最大的矩形
// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/

func largestRectangleArea(heights []int) int {
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
