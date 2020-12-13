package question_581_590

// 590. N叉树的后序遍历
// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal
// Topics: 树

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	for _, children := range root.Children {
		res = append(res, postorder(children)...)
	}
	res = append(res, root.Val)
	return res
}
