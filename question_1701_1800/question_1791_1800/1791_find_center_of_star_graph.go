package question_1791_1800

// 1791. 找出星型图的中心节点
// https://leetcode-cn.com/problems/find-center-of-star-graph/
// Topics: 图

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	} else {
		return edges[0][1]
	}
}
