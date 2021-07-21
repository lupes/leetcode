package question_861_870

import (
	"sort"
)

// 870. 优势洗牌
// https://leetcode-cn.com/problems/advantage-shuffle
// Topics: 贪心算法 数组

func advantageCount(A []int, B []int) []int {
	sort.Ints(A)
	var res = make([]int, len(A))
	for i := range res {
		res[i] = -1
	}
	for _, a := range A {
		var i, max int
		for b := 0; b < len(B); b++ {
			if res[b] == -1 && a > B[b] {
				i = b
				break
			} else if res[b] == -1 && B[b] > max {
				i, max = b, B[b]
			}
		}
		res[i] = a
	}
	return res
}
