package question_1651_1660

// 1653. 使字符串平衡的最少删除次数
// https://leetcode-cn.com/problems/minimum-deletions-to-make-string-balanced/
// Topics: 贪心算法 字符串

func minimumDeletions(s string) int {
	var flag = [2]int{}
	for _, c := range s {
		flag[c-'a']++
	}
	if flag[0] == 0 || flag[1] == 0 {
		return 0
	}

	var a, min, now = flag[0], flag[1], 0
	for _, c := range s {
		if c == 'a' {
			a--
		} else if c == 'b' {
			if now+a < min {
				min = now + a
			}
			now += 1
		}
	}
	return min
}
