package question_171_180

// 173. 二叉搜索树迭代器
// https://leetcode-cn.com/problems/binary-search-tree-iterator
// Topics: 栈 树 设计

import (
	. "github.com/lupes/leetcode/common"
)

type BSTIterator struct {
	now  int
	data []int
}

func BSTIteratorHelper(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	BSTIteratorHelper(res, root.Left)
	*res = append(*res, root.Val)
	BSTIteratorHelper(res, root.Right)
}

func Constructor(root *TreeNode) BSTIterator {
	var res []int
	BSTIteratorHelper(&res, root)
	return BSTIterator{
		now:  -1,
		data: res,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.now++
	return this.data[this.now]
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.now < len(this.data)-1
}
