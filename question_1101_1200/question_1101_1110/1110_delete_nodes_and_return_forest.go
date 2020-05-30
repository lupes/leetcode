package question_1101_1110

// 1110. 删点成林
// https://leetcode-cn.com/problems/delete-nodes-and-return-forest
// Topics: 树 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var m = make(map[int]struct{}, len(to_delete))
	for _, n := range to_delete {
		m[n] = struct{}{}
	}
	if root == nil {
		return nil
	}
	_, res := delNodesHelper(root, m, false)
	return res
}

func delNodesHelper(root *TreeNode, del map[int]struct{}, f bool) (*TreeNode, []*TreeNode) {
	if root == nil {
		return nil, nil
	}
	var res, res1, res2 []*TreeNode
	var n = false
	if _, ok := del[root.Val]; !ok {
		n = true
		if !f {
			res = append(res, root)
		}
	}
	root.Left, res1 = delNodesHelper(root.Left, del, n)
	root.Right, res2 = delNodesHelper(root.Right, del, n)
	res = append(res, res1...)
	res = append(res, res2...)
	if n {
		return nil, res
	} else {
		return root, res
	}
}
