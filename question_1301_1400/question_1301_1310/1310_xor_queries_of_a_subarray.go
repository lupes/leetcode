package question_1301_1310

// 1310. 子数组异或查询
// https://leetcode-cn.com/problems/xor-queries-of-a-subarray/
// Topics: 位运算

func xorQueries(arr []int, queries [][]int) []int {
	var pre, res = make([]int, len(arr)+1), make([]int, len(queries))
	for i := range arr {
		pre[i+1] = pre[i] ^ arr[i]
	}
	for i, query := range queries {
		res[i] = pre[query[1]+1] ^ pre[query[0]]
	}
	return res
}
