package question_111_120

// 111. 二叉树的最小深度
// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Topics: 树 深度优先搜索 广度优先搜索

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if leftDepth == 0 {
		return rightDepth + 1
	}
	if rightDepth == 0 {
		return leftDepth + 1
	}
	min := leftDepth
	if rightDepth < leftDepth {
		min = rightDepth
	}
	return min + 1
}
