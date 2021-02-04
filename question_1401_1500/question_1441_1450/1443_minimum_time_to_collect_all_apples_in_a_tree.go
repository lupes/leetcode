package question_1441_1450

// 1443. 收集树上所有苹果的最少时间
// https://leetcode-cn.com/problems/minimum-time-to-collect-all-apples-in-a-tree/
// Topics: 树 深度优先搜索

type Node struct {
	Val   int
	child []*Node
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	var nodes = make([]Node, n)
	var flag = make([]bool, n)
	for i := range nodes {
		nodes[i].Val = i
	}

	for _, edge := range edges {
		nodes[edge[0]].child = append(nodes[edge[0]].child, &nodes[edge[1]])
		nodes[edge[1]].child = append(nodes[edge[1]].child, &nodes[edge[0]])
	}

	times := 0
	flag[0] = true
	for _, item := range nodes[0].child {
		flag[item.Val] = true
		times += minTimeHelper(item, hasApple, flag)
	}

	return times
}

func minTimeHelper(root *Node, hasApple []bool, flag []bool) (times int) {
	if root == nil {
		return 0
	}

	for _, item := range root.child {
		if !flag[item.Val] {
			flag[item.Val] = true
			times += minTimeHelper(item, hasApple, flag)
		}
	}

	if times > 0 || (times == 0 && hasApple[root.Val]) {
		times += 2
	}

	return times
}
