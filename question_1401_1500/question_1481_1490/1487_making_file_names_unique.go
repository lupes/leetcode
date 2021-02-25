package question_1481_1490

// 1487. 保证文件名唯一
// https://leetcode-cn.com/problems/making-file-names-unique/
// Topics: 哈希表 字符串

import (
	"fmt"
)

func getFolderNames(names []string) []string {
	var flag = make(map[string]int)
	var res = make([]string, 0)
	for _, name := range names {
		if n, ok := flag[name]; !ok {
			flag[name] = 0
			res = append(res, name)
		} else {
			for i := n + 1; ; i++ {
				tmp := fmt.Sprintf("%s(%d)", name, i)
				if _, ok := flag[tmp]; !ok {
					flag[tmp] = 0
					flag[name] = i
					res = append(res, tmp)
					break
				}
			}
		}
	}

	return res
}
