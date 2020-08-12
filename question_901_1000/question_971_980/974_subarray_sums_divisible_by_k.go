package question_971_980

// 974. 和可被 K 整除的子数组
// https://leetcode-cn.com/problems/subarray-sums-divisible-by-k
// Topics: 数组 哈希表

func subarraysDivByK(A []int, K int) int {
	var res, sum, flag = 0, 0, make(map[int]int)
	flag[0] = 1
	for i := 0; i < len(A); i++ {
		sum += A[i]
		mod := (sum%K + K) % K
		res += flag[mod]
		flag[mod]++
	}
	return res
}
