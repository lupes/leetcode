package question_381_390

// 386. 字典序排数
// https://leetcode-cn.com/problems/lexicographical-numbers
// Topics:

func lexicalOrder(n int) []int {
	var res = make([]int, 0, n)
	for i := 1; i < 10; i++ {
		lexicalOrderDFS(i, n, &res)
	}
	return res
}

func lexicalOrderDFS(start, n int, res *[]int) {
	if start > n {
		return
	}
	*res = append(*res, start)
	for i := 0; i < 10; i++ {
		lexicalOrderDFS(start*10+i, n, res)
	}
}
