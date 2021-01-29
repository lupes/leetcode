package question_1421_1430

// 1424. 对角线遍历 II
// https://leetcode-cn.com/problems/diagonal-traverse-ii/
// Topics: 排序 数组

func findDiagonalOrder(nums [][]int) []int {
	var tmp = make(map[int][]int)
	for i, row := range nums {
		for j, n := range row {
			tmp[i+j] = append(tmp[i+j], n)
		}
	}

	var res []int
	for i := 0; ; i++ {
		arr, ok := tmp[i]
		if !ok {
			return res
		}
		for i := len(arr) - 1; i >= 0; i-- {
			res = append(res, arr[i])
		}
	}
}
