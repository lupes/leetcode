package question_41_50

import (
	"sort"
	"strings"
)

// 49. 字母异位词分组
// https://leetcode-cn.com/problems/group-anagrams

func groupAnagrams(strs []string) [][]string {
	var tmp = make(map[string]int)
	var res [][]string
	for _, str := range strs {
		arr := strings.Split(str, "")
		sort.Strings(arr)
		s := strings.Join(arr, "")
		value, ok := tmp[s]
		if !ok {
			res = append(res, []string{str})
			tmp[s] = len(res) - 1
		} else {
			res[value] = append(res[value], str)
		}
	}
	return res
}
