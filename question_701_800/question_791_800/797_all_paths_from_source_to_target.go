package question_791_800

// 797. 所有可能的路径
// https://leetcode-cn.com/problems/all-paths-from-source-to-target
// Topics: 深度优先搜索 广度优先搜索 图 回溯

func allPathsSourceTarget(graph [][]int) [][]int {
	return allPathsSourceTargetHelper(graph, 0, len(graph)-1)
}

func allPathsSourceTargetHelper(graph [][]int, src, dst int) [][]int {
	if src == dst {
		return [][]int{{dst}}
	}
	var res [][]int
	for _, n := range graph[src] {
		tmp := allPathsSourceTargetHelper(graph, n, dst)
		for _, f := range tmp {
			res = append(res, append([]int{src}, f...))
		}
	}
	return res
}
