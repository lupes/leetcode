package question_1231_1240

// 1237. 找出给定方程的正整数解
// https://leetcode-cn.com/problems/find-positive-integer-solution-for-a-given-equation
// Topics: 数学 二分查找

func findSolution(customFunction func(int, int) int, z int) [][]int {
	var res, t = make([][]int, 0), z
	for i := 1; i <= z; i++ {
		for j := 1; j <= t; j++ {
			if customFunction(i, j) == z {
				res, t = append(res, []int{i, j}), j
			}
		}
	}
	return res
}
