package question_1521_1530

// 1524. 和为奇数的子数组数目
// https://leetcode-cn.com/problems/number-of-sub-arrays-with-odd-sum/
// Topics: 数学 数组

func numOfSubarrays(arr []int) int {
	var sum, odd, even = 0, 0, 1
	for _, n := range arr {
		sum += n
		if sum%2 == 0 {
			even += 1
		} else {
			odd += 1
		}
	}
	return (even * odd) % (1e9 + 7)
}
