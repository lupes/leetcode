package question_1641_1650

// 1647. 字符频次唯一的最小删除次数
// https://leetcode-cn.com/problems/minimum-deletions-to-make-character-frequencies-unique/
// Topics: 贪心算法 排序

import (
	"sort"
)

func minDeletions(s string) int {
	var flag = make([]int, 26)
	for _, c := range s {
		flag[c-'a']++
	}

	sort.Slice(flag, func(i, j int) bool {
		return flag[i] > flag[j]
	})

	var res = 0
	for i := 1; i < 26 && flag[i] > 0; i++ {
		if flag[i] >= flag[i-1] {
			d := flag[i] - flag[i-1] + 1
			if flag[i]-d < 0 {
				d = flag[i]
			}
			res += d
			flag[i] -= d
		}
	}
	return res
}
