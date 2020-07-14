package question_421_430

// 429. N叉树的层序遍历
// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal
// Topics: 树 广度优先搜索

type Node5 struct {
	Val      int
	Children []*Node5
}

func levelOrder(root *Node5) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var now = []*Node5{root}
	for len(now) > 0 {
		var next []*Node5
		var tmp []int
		for _, n := range now {
			tmp = append(tmp, n.Val)
			for _, c := range n.Children {
				next = append(next, c)
			}
		}
		now = next
		res = append(res, tmp)
	}
	return res
}
