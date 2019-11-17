package question_111_120

// 112. 路径总和
// https://leetcode-cn.com/problems/path-sum/
// Topics: 树 深度优先搜索

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
