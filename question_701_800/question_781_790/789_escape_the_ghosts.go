package question_781_790

// 789. 逃脱阻碍者
// https://leetcode-cn.com/problems/escape-the-ghosts
// Topics: 数学

func escapeGhosts(ghosts [][]int, target []int) bool {
	t := escapeGhostsHelper([]int{0, 0}, target)
	if t == 0 {
		return true
	}
	for _, ghost := range ghosts {
		a := escapeGhostsHelper(ghost, target)
		if a <= t {
			return false
		}
	}
	return true
}

func escapeGhostsHelper(src, dest []int) int {
	var res = 0
	if t := src[0] - dest[0]; t >= 0 {
		res += t
	} else {
		res += -t
	}

	if t := src[1] - dest[1]; t >= 0 {
		res += t
	} else {
		res += -t
	}
	return res
}
