package question_71_80

// 76. 最小覆盖子串
// https://leetcode-cn.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	// step1 将t转换为map结构 切片长度为2，第一个存储在t的出现次数，第二个存储在当前窗口出现的次数
	var flags = make(map[byte][]int)
	for i := range t {
		if _, ok := flags[t[i]]; !ok {
			flags[t[i]] = make([]int, 2)
		}
		flags[t[i]][0]++
	}

	// step2 移动左右指针，首先移动右指针直到满足题目要求，然后移动左指针，直到不满足要求
	left, right, res, min, required := 0, 0, "", len(s)+1, len(flags)
	for right <= len(s) {
		if required == 0 {
			if flag, ok := flags[s[left]]; ok {
				flag[1]--
				if flag[1] == flag[0]-1 {
					required++
				}
			}
			left++
		} else {
			if right == len(s) {
				break
			}
			if flag, ok := flags[s[right]]; ok {
				flag[1]++
				if flag[0] == flag[1] {
					required--
				}
			}
			right++
		}
		if required == 0 && right-left < min {
			min, res = right-left, s[left:right]
		}
	}
	return res
}
