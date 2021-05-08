package question_1681_1690

// 1685. 有序数组中差绝对值之和
// https://leetcode-cn.com/problems/sum-of-absolute-differences-in-a-sorted-array/
// Topics: 贪心算法 数组

func getSumAbsoluteDifferences(nums []int) []int {
	var res = make([]int, len(nums))
	for _, n := range nums {
		res[0] += n - nums[0]
	}
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] + (nums[i]-nums[i-1])*(2*i-len(nums))
	}
	return res
}
