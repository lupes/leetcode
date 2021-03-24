package question_1551_1560

// 1557. 可以到达所有点的最少点数目
// https://leetcode-cn.com/problems/minimum-number-of-vertices-to-reach-all-nodes/
// Topics: 图

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	var flag = make([]bool, n)
	for _, edge := range edges {
		flag[edge[1]] = true
	}
	var res []int
	for i, t := range flag {
		if !t {
			res = append(res, i)
		}
	}
	return res
}
