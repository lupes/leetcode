package question_131_140

// 133. 克隆图
// https://leetcode-cn.com/problems/clone-graph
// Topics: 深度优先搜索 广度优先搜索 图

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	return cloneGraphHelper(node, make(map[int]*Node))
}

func cloneGraphHelper(node *Node, flag map[int]*Node) *Node {
	flag[node.Val] = &Node{Val: node.Val}
	for _, n := range node.Neighbors {
		t, ok := flag[n.Val]
		if !ok {
			t = cloneGraphHelper(n, flag)
		}
		flag[node.Val].Neighbors = append(flag[node.Val].Neighbors, t)
	}
	return flag[node.Val]
}
