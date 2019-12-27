package question_641_650

// 643. 子数组最大平均数 I
// https://leetcode-cn.com/problems/maximum-average-subarray-i
// Topics: 数组

func findMaxAverage(nums []int, k int) float64 {
	sum, i := 0, 0
	for ; i < k; i++ {
		sum += nums[i]
	}
	avg := float64(sum) / float64(k)
	for ; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		tmp := float64(sum) / float64(k)
		if tmp > avg {
			avg = tmp
		}
	}
	return avg
}
