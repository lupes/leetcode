package question_371_380

// 376. 摆动序列
// https://leetcode-cn.com/problems/wiggle-subsequence
// Topics: 贪心算法 动态规划

func wiggleMaxLength(nums []int) int {
	r1 := wiggleMaxLengthHelper(nums, true)
	r2 := wiggleMaxLengthHelper(nums, false)
	if r1 > r2 {
		return r1
	}
	return r2
}

func wiggleMaxLengthHelper(nums []int, flag bool) int {
	var res, tmp = 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if (flag && (nums[i] > tmp)) || (!flag && nums[i] < tmp) {
			res, tmp, flag = res+1, nums[i], !flag
		} else {
			tmp = nums[i]
		}
	}
	return res
}
