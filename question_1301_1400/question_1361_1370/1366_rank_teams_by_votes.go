package question_1361_1370

import (
	"sort"
)

// 1366. 通过投票对团队排名
// https://leetcode-cn.com/problems/rank-teams-by-votes/
// Topics: 排序 数组

func rankTeams(votes []string) string {
	var flag = make(map[int32][]int32)
	for _, vote := range votes {
		for i, c := range vote {
			if _, ok := flag[c]; !ok {
				flag[c] = make([]int32, len(vote)+1)
				flag[c][0] = c
			}
			flag[c][i+1]++
		}
	}
	var tmp = make([][]int32, 0, len(flag))
	for _, v := range flag {
		tmp = append(tmp, v)
	}
	sort.Slice(tmp, func(i, j int) bool {
		for k := 1; k < len(tmp[0]); k++ {
			if tmp[i][k] > tmp[j][k] {
				return true
			} else if tmp[i][k] < tmp[j][k] {
				return false
			}
		}
		return tmp[i][0] < tmp[j][0]
	})
	var res = make([]byte, len(tmp))
	for i, t := range tmp {
		res[i] = byte(t[0])
	}
	return string(res)
}
