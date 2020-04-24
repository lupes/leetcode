package question_1001_1010

// 1002. 查找常用字符
// https://leetcode-cn.com/problems/find-common-characters
// Topics: 数组 哈希表

func commonChars(A []string) []string {
	var flag = make([][26]int, len(A))
	for i, a := range A {
		for _, c := range a {
			flag[i][c-'a']++
		}
	}
	var res []string
	for i := 0; i < 26; i++ {
		n := 100
		for _, f := range flag {
			if f[i] < n {
				n = f[i]
			}
		}
		for j := 0; j < n; j++ {
			res = append(res, string(i+'a'))
		}
	}
	return res
}
