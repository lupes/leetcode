package question_111_120

// 119. 杨辉三角 II
// https://leetcode-cn.com/problems/pascals-triangle-ii/
// Topics: 数组

func getRow(rowIndex int) []int {
	var res []int
	for i := 0; i <= rowIndex; i++ {
		var t = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				t[j] = 1
			} else {
				t[j] = res[j-1] + res[j]
			}
		}
		res = t
	}
	return res
}
