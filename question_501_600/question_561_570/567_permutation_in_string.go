package question_561_570

// 567. 字符串的排列
// https://leetcode-cn.com/problems/permutation-in-string
// Topics: 双指针 None

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var flag = make(map[uint8]int)
	for i := range s1 {
		flag[s1[i]]++
	}
	sum, l, zero := len(flag), len(s1), 0
	for left, right := 0, 0; len(s2) > right; {
		if (right - left) < l {
			flag[s2[right]]--
			if flag[s2[right]] == 0 {
				zero++
			} else if flag[s2[right]] == -1 {
				zero--
			}
			right++
		} else {
			flag[s2[left]]++
			if flag[s2[left]] == 0 {
				zero++
			} else if flag[s2[left]] == 1 {
				zero--
			}
			left++
		}
		if zero == sum {
			return true
		}
	}
	return false
}
