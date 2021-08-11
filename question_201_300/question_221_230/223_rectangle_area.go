package question_221_230

// 223. 矩形面积
// https://leetcode-cn.com/problems/rectangle-area
// Topics: 数学

func computeAreaLeft(a1, a2, b1, b2 int) int {
	if a1 >= b1 && a1 < b2 {
		return a1
	} else if b1 >= a1 && b1 < a2 {
		return b1
	}
	return 0
}

func computeAreaRight(a1, a2, b1, b2 int) int {
	if a2 > b1 && a2 <= b2 {
		return a2
	} else if b2 > a1 && b2 <= a2 {
		return b2
	}
	return 0
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	x1 := computeAreaLeft(ax1, ax2, bx1, bx2)
	y1 := computeAreaLeft(ay1, ay2, by1, by2)
	x2 := computeAreaRight(ax1, ax2, bx1, bx2)
	y2 := computeAreaRight(ay1, ay2, by1, by2)

	return (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1) - (x2-x1)*(y2-y1)
}
