package question_471_480

// 477. 汉明距离总和
// https://leetcode-cn.com/problems/total-hamming-distance
// Topics: 位运算

func totalHammingDistance(nums []int) int {
	var res = 0
	for i := 0; i < 31; i++ {
		var tmp [2]int
		for i, n := range nums {
			tmp[n&1]++
			nums[i] >>= 1
		}
		res += tmp[0] * tmp[1]
	}
	return res
}
