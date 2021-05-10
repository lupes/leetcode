package question_1671_1680

// 1679. K 和数对的最大数目
// https://leetcode-cn.com/problems/max-number-of-k-sum-pairs/
// Topics: 哈希表

func maxOperations(nums []int, k int) int {
	var flag, res = make(map[int]int), 0
	for _, num := range nums {
		flag[num] += 1
		if k-num > 0 {
			t := k - num
			if (t == num && flag[num] > 1) || (t != num && flag[t] > 0 && flag[num] > 0) {
				res += 1
				flag[num] -= 1
				flag[t] -= 1
			}
		}
	}
	return res
}
