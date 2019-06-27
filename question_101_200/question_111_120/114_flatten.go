package question_111_120

// 114. 二叉树展开为链表
// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	tmp := root.Right
	root.Right = root.Left
	root.Left = nil
	next := root
	for next.Right != nil {
		next = next.Right
	}
	next.Right = tmp
}
