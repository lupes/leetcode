package question_1591_1600

// 1593. 拆分字符串使唯一子字符串的数目最大
// https://leetcode-cn.com/problems/split-a-string-into-the-max-number-of-unique-substrings/
// 回溯算法

func maxUniqueSplit(s string) int {
	return maxUniqueSplitHelper(s, make(map[string]struct{}), 0)
}

func maxUniqueSplitHelper(s string, flag map[string]struct{}, n int) int {
	if len(s) == 0 {
		return n
	}
	var res = 0
	for i := 1; i <= len(s); i++ {
		if _, ok := flag[s[0:i]]; !ok {
			flag[s[0:i]] = struct{}{}
			tmp := maxUniqueSplitHelper(s[i:], flag, n+1)
			if tmp > res {
				res = tmp
			}
			delete(flag, s[0:i])
		}
	}
	return res
}
