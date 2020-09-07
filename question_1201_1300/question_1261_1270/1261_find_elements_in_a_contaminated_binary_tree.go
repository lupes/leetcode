package question_1261_1270

import (
	. "github.com/lupes/leetcode/common"
)

// 1261. 在受污染的二叉树中查找元素
// https://leetcode-cn.com/problems/find-elements-in-a-contaminated-binary-tree/
// Topics: 树 哈希表

type FindElements struct {
	flag map[int]struct{}
}

func FindElementsHelper(root *TreeNode, flag map[int]struct{}) {
	if root.Left != nil {
		root.Left.Val = root.Val*2 + 1
		flag[root.Left.Val] = struct{}{}
		FindElementsHelper(root.Left, flag)
	}
	if root.Right != nil {
		root.Right.Val = root.Val*2 + 2
		flag[root.Right.Val] = struct{}{}
		FindElementsHelper(root.Right, flag)
	}
}

func Constructor(root *TreeNode) FindElements {
	root.Val = 0
	var flag = make(map[int]struct{})
	flag[0] = struct{}{}
	FindElementsHelper(root, flag)
	return FindElements{flag: flag}
}

func (this *FindElements) Find(target int) bool {
	_, ok := this.flag[target]
	return ok
}
