package question_561_570

// 566. 重塑矩阵
// https://leetcode-cn.com/problems/reshape-the-matrix/

func matrixReshape(nums [][]int, r int, c int) [][]int {
	nr, nc := len(nums), len(nums[0])
	if nr*nc != r*c {
		return nums
	}
	var res = make([][]int, r)
	for i := 0; i < nr*nc; i++ {
		res[i/c] = append(res[i/c], nums[i/nc][i%nc])
	}
	return res
}
