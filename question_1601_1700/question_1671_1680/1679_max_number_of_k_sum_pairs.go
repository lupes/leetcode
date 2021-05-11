package question_1671_1680

// 1679. K 和数对的最大数目
// https://leetcode-cn.com/problems/max-number-of-k-sum-pairs/
// Topics: 哈希表

func maxOperations(nums []int, k int) int {
	var flag, tmp, res = make(map[int]int), 0, 0
	for _, num := range nums {
		tmp = k - num
		if tmp > 0 {
			if flag[tmp] > 0 {
				res += 1
				flag[tmp] -= 1
			} else {
				flag[num] += 1
			}
		}
	}
	return res
}
