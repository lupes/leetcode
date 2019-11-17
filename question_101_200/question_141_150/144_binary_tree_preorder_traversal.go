package question_141_150

// 144. 二叉树的前序遍历
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
// Topics: 栈 树

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack = []*TreeNode{root}
	var res []int
	var node *TreeNode
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		res = append(res, node.Val)

		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
