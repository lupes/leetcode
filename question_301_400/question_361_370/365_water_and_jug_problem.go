package question_361_370

// 365. 水壶问题
// https://leetcode-cn.com/problems/water-and-jug-problem
// Topics: 数学

func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}
	if x > y {
		x, y = y, x
	}
	for y%x != 0 {
		y, x = x, y%x
	}
	return z%x == 0
}
