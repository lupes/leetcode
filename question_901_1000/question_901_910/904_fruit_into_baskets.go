package question_901_910

// 904. 水果成篮
// https://leetcode-cn.com/problems/fruit-into-baskets
// Topics: 双指针

func totalFruit(tree []int) int {
	var flag = make(map[int]int)
	l, left, right, max := len(tree), 0, 0, 0
	for right >= left && l > right {
		if len(flag) < 3 && right < l {
			flag[tree[right]]++
			right++
		} else {
			flag[tree[left]]--
			if flag[tree[left]] == 0 {
				delete(flag, tree[left])
			}
			left++
		}
		if len(flag) < 3 {
			if right-left > max {
				max = right - left
			}
		}
	}
	return max
}
