package question_651_660

import (
	"encoding/json"

	. "github.com/lupes/leetcode/common"
)

// 652. 寻找重复的子树
// https://leetcode-cn.com/problems/find-duplicate-subtrees
// Topics: 树

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var flag = make(map[string]*TreeNode)
	var resFlag = make(map[*TreeNode]struct{})
	var res = make([]*TreeNode, 0)
	findDuplicateSubtreesHelper(root, flag, resFlag, &res)
	return res
}

func findDuplicateSubtreesHelper(root *TreeNode, flag map[string]*TreeNode, resFlag map[*TreeNode]struct{}, res *[]*TreeNode) {
	if root == nil {
		return
	}
	buffer, _ := json.Marshal(root)
	node, ok := flag[string(buffer)]
	if ok {
		if _, ok := resFlag[node]; !ok {
			resFlag[node] = struct{}{}
			*res = append(*res, node)
		}
	} else {
		flag[string(buffer)] = root
	}
	findDuplicateSubtreesHelper(root.Left, flag, resFlag, res)
	findDuplicateSubtreesHelper(root.Right, flag, resFlag, res)
}
