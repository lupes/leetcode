package question_551_560

// 559. N叉树的最大深度
// https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree
// Topics: 树 深度优先搜索 广度优先搜索

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	var max, tmp = 0, 0
	if root == nil {
		return 0
	}

	for _, ch := range root.Children {
		tmp = maxDepth(ch)
		if tmp > max {
			max = tmp
		}
	}
	return max + 1
}
