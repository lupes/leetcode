package question_851_860

// 859. 亲密字符串
// https://leetcode-cn.com/problems/buddy-strings
// Topics: 字符串 哈希表

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	var a, b, num int
	var flag = make([]int, 26)
	for i := range A {
		if A[i] != B[i] {
			num++
			if num == 1 {
				a = i
			} else if num == 2 {
				b = i
			} else if num == 3 {
				return false
			}
		} else {
			flag[A[i]-'a']++
		}
	}
	if num == 0 {
		for _, t := range flag {
			if t > 1 {
				return true
			}
		}
	}
	return num == 2 && A[a] == B[b] && A[b] == B[a]
}
