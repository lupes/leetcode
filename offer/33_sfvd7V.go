package offer

// 剑指 Offer II 033. 变位词组
// https://leetcode-cn.com/problems/sfvd7V/
// 哈希表 字符串 排序

func groupAnagrams(strs []string) [][]string {
	var flag = make(map[string][]string)
	for _, str := range strs {
		key := groupAnagramsHash(str)
		flag[key] = append(flag[key], str)
	}
	var res [][]string
	for _, values := range flag {
		res = append(res, values)
	}
	return res
}

func groupAnagramsHash(str string) string {
	var flag = make([]byte, 26)
	for _, c := range str {
		flag[c-'a']++
	}
	return string(flag)
}
