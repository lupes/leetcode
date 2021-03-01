package question_1491_1500

// 1497. 检查数组对是否可以被 k 整除
// https://leetcode-cn.com/problems/check-if-array-pairs-are-divisible-by-k/
// Topics: 贪心算法 数组 数学

func canArrange(arr []int, k int) bool {
	var flag, a, b = make(map[int]int, 0), 0, 0
	for _, n := range arr {
		a, b = n%k, 0
		if a < 0 {
			a, b = k+a, k-(k+a)
		} else if a > 0 {
			b = k - a
		}

		if flag[b] == 1 {
			delete(flag, b)
		} else if flag[b] > 1 {
			flag[b]--
		} else {
			flag[a]++
		}
	}
	return len(flag) == 0
}
