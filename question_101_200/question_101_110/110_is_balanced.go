package question_101_110

// 110. 平衡二叉树
// https://leetcode-cn.com/problems/balanced-binary-tree/submissions/

func isBalanced(root *TreeNode) bool {
	_, flag := maxDepth2(root)
	return flag
}

func maxDepth2(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, leftFlag := maxDepth2(root.Left)
	rightDepth, rightFlag := maxDepth2(root.Right)
	max := leftDepth
	if rightDepth > leftDepth {
		max = rightDepth
	}
	if leftFlag && rightFlag && (rightDepth+1 == leftDepth || leftDepth+1 == rightDepth || leftDepth == rightDepth) {
		return max + 1, true
	} else {
		return max + 1, false
	}
}
