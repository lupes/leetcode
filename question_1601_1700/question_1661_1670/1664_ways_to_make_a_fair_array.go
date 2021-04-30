package question_1661_1670

// 1664. 生成平衡数组的方案数
// https://leetcode-cn.com/problems/ways-to-make-a-fair-array/
// Topics: 贪心算法 动态规划

func waysToMakeFair(nums []int) int {
	var lSums, rSums, l = make([]int, len(nums)+2), make([]int, len(nums)+2), len(nums)
	for i := 0; i < l; i++ {
		lSums[i+2] += nums[i] + lSums[i]
		rSums[l-i-1] += rSums[l+1-i] + nums[l-i-1]
	}

	var r = 0
	for i := 0; i < l; i++ {
		if lSums[i]+rSums[i+1] == lSums[i+1]+rSums[i+2] {
			r++
		}
	}
	return r
}
