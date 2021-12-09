package question_681_690

// 689. 三个无重叠子数组的最大和
// https://leetcode-cn.com/problems/maximum-sum-of-3-non-overlapping-subarrays
// Topics: 数组 动态规划

func maxSumOfThreeSubarrays(nums []int, k int) (ans []int) {
	var k1, k2, k3 = k, k * 2, k * 3
	var sum1, maxSum1, maxSum1Idx int
	var sum2, maxSum2, maxSum2Idx1, maxSum2Idx2 int
	var sum3, maxSum3 int
	for i := k2; i < len(nums); i++ {
		sum1 += nums[i-k2]
		sum2 += nums[i-k1]
		sum3 += nums[i]
		if i < k3-1 {
			continue
		}

		if sum1 > maxSum1 {
			maxSum1 = sum1
			maxSum1Idx = i - k3 + 1
		}
		if maxSum1+sum2 > maxSum2 {
			maxSum2 = maxSum1 + sum2
			maxSum2Idx1, maxSum2Idx2 = maxSum1Idx, i-k2+1
		}
		if maxSum2+sum3 > maxSum3 {
			maxSum3 = maxSum2 + sum3
			ans = []int{maxSum2Idx1, maxSum2Idx2, i - k1 + 1}
		}
		sum1 -= nums[i-k3+1]
		sum2 -= nums[i-k2+1]
		sum3 -= nums[i-k1+1]
	}
	return
}
