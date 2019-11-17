package question_331_340

import (
	"sort"
)

// 332. 重新安排行程
// https://leetcode-cn.com/problems/reconstruct-itinerary/
// Topics: 深度优先搜索 图

func findItinerary(tickets [][]string) []string {
	var m = make(map[string][]string, 0)
	for _, ticket := range tickets {
		m[ticket[0]] = append(m[ticket[0]], ticket[1])
	}

	for key, vs := range m {
		sort.Strings(vs)
		m[key] = vs
	}

	var res = []string{"JFK"}
	for {
		key := res[len(res)-1]
		vs := m[key]
		if len(vs) == 0 {
			if len(m["JFK"]) != 0 {
				res = append(res, "JFK")
				continue
			} else {
				return res
			}
		}
		res = append(res, vs[0])
		m[key] = vs[1:]
	}
}
