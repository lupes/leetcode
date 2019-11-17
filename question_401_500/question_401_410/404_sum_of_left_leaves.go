package question_401_410

// 404. 左叶子之和
// https://leetcode-cn.com/problems/sum-of-left-leaves/
// Topics: 树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
	}

	if root.Left != nil {
		res += sumOfLeftLeaves(root.Left)
	}
	if root.Right != nil {
		res += sumOfLeftLeaves(root.Right)
	}
	return res
}
