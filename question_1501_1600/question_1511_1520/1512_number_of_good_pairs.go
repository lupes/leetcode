package question_1511_1520

// 1512. 好数对的数目
// https://leetcode-cn.com/problems/number-of-good-pairs/
// Topics: 数组 哈希表 数学

func numIdenticalPairs(nums []int) int {
	var flag = make(map[int]int, 0)
	for _, n := range nums {
		flag[n]++
	}

	var res = 0
	for _, num := range flag {
		res += num * (num - 1) / 2
	}
	return res
}
