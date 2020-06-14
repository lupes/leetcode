package question_1231_1240

import (
	"sort"
)

// 1233. 删除子文件夹
// https://leetcode-cn.com/problems/remove-sub-folders-from-the-filesystem
// Topics: 数组 字符串

func in(flag map[string]struct{}, f string) bool {
	for i, c := range f {
		if c == '/' {
			if _, ok := flag[f[:i+1]]; ok {
				return true
			}
		}
	}
	return false
}

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	var ok, res, flag = false, make([]string, 0), make(map[string]struct{})
	for _, f := range folder {
		if ok = in(flag, f+"/"); !ok {
			flag[f+"/"] = struct{}{}
		}
	}
	for k, _ := range flag {
		res = append(res, k[:len(k)-1])
	}
	return res
}
