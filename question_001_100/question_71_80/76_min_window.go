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

	// step2 设置快慢指针，slow为0，fast为从0开始满足条件的第一个子串，如果某个字符不满足最小出现次数返回空字符串
	fast, slow := 0, 0
	for k, flag := range flags {
		if flag[0] <= flag[1] {
			continue
		}
		for ; fast < len(s); fast++ {
			if flag, ok := flags[s[fast]]; ok {
				flag[1]++
				if k == s[fast] && flag[0] <= flag[1] {
					fast++
					break
				}
			}
		}
		if flag[0] > flag[1] {
			return ""
		}
	}

	// step3 移动快慢指针，首先移动慢指针直到某个字符出现次数少于指定次数，然后移动快指针，直到该字符满足出现次数
	next, min, res := byte(0), fast-slow, s[slow:fast]
	for slow < len(s) && fast <= len(s) {
		if next == 0 {
			if flag, ok := flags[s[slow]]; ok {
				flag[1]--
				if flag[1] < flag[0] {
					next = s[slow]
				}
			}
			slow++
		} else {
			if fast == len(s) {
				break
			}
			if flag, ok := flags[s[fast]]; ok {
				flag[1]++
			}
			if s[fast] == next {
				next = 0
			}
			fast++
		}
		if next == 0 && fast-slow < min {
			min, res = fast-slow, s[slow:fast]
		}
	}
	return res
}
