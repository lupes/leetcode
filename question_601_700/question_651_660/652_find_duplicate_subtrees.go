package question_651_660

import (
	"fmt"

	. "github.com/lupes/leetcode/common"
)

// 652. 寻找重复的子树
// https://leetcode-cn.com/problems/find-duplicate-subtrees
// Topics: 树

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var flags = make(map[string]bool)
	var res = make([]*TreeNode, 0)
	findDuplicateSubtreesHelper(root, flags, &res)
	return res
}

func findDuplicateSubtreesHelper(root *TreeNode, flags map[string]bool, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}
	left := findDuplicateSubtreesHelper(root.Left, flags, res)
	right := findDuplicateSubtreesHelper(root.Right, flags, res)
	str := fmt.Sprintf("%d,%s,%s", root.Val, left, right)
	flag, ok := flags[str]
	if ok {
		if !flag {
			flags[str] = true
			*res = append(*res, root)
		}
	} else {
		flags[str] = false
	}
	return str
}
