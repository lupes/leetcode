package question_851_860

// 851. 喧闹和富有
// https://leetcode-cn.com/problems/loud-and-rich
// Topics: 深度优先搜索

type node struct {
	p    int
	rich []*node
}

func loudAndRich(richer [][]int, quiet []int) []int {
	var flag = make(map[int]*node)
	var res = make([]int, len(quiet))
	for i := range quiet {
		res[i] = -1
		flag[i] = &node{p: i}
	}

	for _, r := range richer {
		flag[r[1]].rich = append(flag[r[1]].rich, flag[r[0]])
	}

	for i := range quiet {
		loudAndRichDfs(flag, quiet, res, i)
	}
	return res
}

func loudAndRichDfs(flag map[int]*node, quiet, res []int, i int) int {
	if res[i] == -1 {
		res[i] = i
		min := quiet[i]
		for _, n := range flag[i].rich {
			tmp := loudAndRichDfs(flag, quiet, res, n.p)
			if tmp < min {
				min = tmp
				res[i] = res[n.p]
			}
		}
	}
	return quiet[res[i]]
}
