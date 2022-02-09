package question_2001_2010

// 2006. 差的绝对值为 K 的数对数目
// https://leetcode-cn.com/problems/count-number-of-pairs-with-absolute-difference-k/
// Topic: 数组 哈希表 计数

func countKDifference(nums []int, k int) int {
	var flag, res = make(map[int]int), 0
	for _, n := range nums {
		res = res + flag[n-k] + flag[n+k]
		flag[n]++
	}
	return res
}
