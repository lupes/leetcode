package question_1241_1250

import (
	"sort"
)

// 1247. 交换字符使得字符串相同
// https://leetcode-cn.com/problems/minimum-swaps-to-make-strings-equal
// Topics: 贪心算法 字符串

func minimumSwap(s1 string, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}
	type node struct {
		c1 byte
		c2 byte
	}
	var nodes []node
	for i := range s1 {
		if s1[i] != s2[i] {
			nodes = append(nodes, node{
				c1: s1[i],
				c2: s2[i],
			})
		}
	}
	if len(nodes)%2 != 0 {
		return -1
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].c1 < nodes[j].c1 || (nodes[i].c1 == nodes[j].c1 && nodes[i].c2 < nodes[j].c2)
	})

	var res = 0
	for i := 0; i < len(nodes); i += 2 {
		if (nodes[i].c1 == 'x' && nodes[i+1].c1 == 'x') || (nodes[i].c1 == 'y' && nodes[i+1].c1 == 'y') {
			res += 1
		} else {
			res += 2
		}
	}

	return res
}
