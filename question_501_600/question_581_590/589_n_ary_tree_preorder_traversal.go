package question_581_590

// 589. N叉树的前序遍历
// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal
// Topics: 树

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	res = append(res, root.Val)
	for _, children := range root.Children {
		res = append(res, preorder(children)...)
	}
	return res
}
