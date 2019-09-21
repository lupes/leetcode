package question_461_470

// 461. 汉明距离
// https://leetcode-cn.com/problems/hamming-distance/

func hammingDistance(x int, y int) int {
	t, r := x^y, 0
	for t > 0 {
		r += t & 1
		t >>= byte(1)
	}
	return r
}
