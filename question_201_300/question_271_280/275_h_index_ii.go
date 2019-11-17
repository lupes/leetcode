package question_271_280

// 275. H指数 II
// https://leetcode-cn.com/problems/h-index-ii/
// Topics: 二分查找

func hIndex2(citations []int) int {
	l := len(citations)
	for i, n := range citations {
		if n >= l-i {
			return l - i
		}
	}
	return 0
}
