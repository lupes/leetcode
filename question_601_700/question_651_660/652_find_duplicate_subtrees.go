package question_651_660

import (
	"fmt"

	. "github.com/lupes/leetcode/common"
)

// 652. 寻找重复的子树
// https://leetcode-cn.com/problems/find-duplicate-subtrees
// Topics: 树

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var flag = make(map[string]struct{})
	var resFlag = make(map[string]struct{})
	var res = make([]*TreeNode, 0)
	findDuplicateSubtreesHelper(root, flag, resFlag, &res)
	return res
}

func findDuplicateSubtreesHelper(root *TreeNode, flag map[string]struct{}, resFlag map[string]struct{}, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}
	left := findDuplicateSubtreesHelper(root.Left, flag, resFlag, res)
	right := findDuplicateSubtreesHelper(root.Right, flag, resFlag, res)
	str := fmt.Sprintf("%d,%s,%s", root.Val, left, right)
	_, ok := flag[str]
	if ok {
		if _, ok := resFlag[str]; !ok {
			resFlag[str] = struct{}{}
			*res = append(*res, root)
		}
	} else {
		flag[str] = struct{}{}
	}
	return str
}
