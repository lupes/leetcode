package question_491_500

import "math"

// 492. 构造矩形
// https://leetcode-cn.com/problems/construct-the-rectangle/

func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
	for ; area%w != 0; w-- {
	}
	return []int{area / w, w}
}
