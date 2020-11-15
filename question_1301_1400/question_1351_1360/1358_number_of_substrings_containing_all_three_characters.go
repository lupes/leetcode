package question_1331_1340

// 1358. 包含所有三种字符的子字符串数目
// https://leetcode-cn.com/problems/number-of-substrings-containing-all-three-characters/
// Topics: 字符串

func numberOfSubstrings(s string) int {
	nums, left, right, flag, res := [3]int{}, 0, 0, true, 0
	for right < len(s) {
		if flag {
			nums[s[right]-'a']++
			flag = false
		}
		if nums[0] > 0 && nums[1] > 0 && nums[2] > 0 {
			res += len(s) - right
			nums[s[left]-'a']--
			left++
		} else {
			right++
			flag = true
		}
	}
	return res
}
