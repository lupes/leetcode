package question_1311_1320

// 1315. 祖父节点值为偶数的节点和
// https://leetcode-cn.com/problems/sum-of-nodes-with-even-valued-grandparent/
// Topics: 树 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func sumEvenGrandparent(root *TreeNode) int {
	return sumEvenGrandparentHelp(root, -1, -1)
}

func sumEvenGrandparentHelp(root *TreeNode, p, pp int) int {
	if root == nil {
		return 0
	}
	var res = 0
	if pp == 0 {
		res += root.Val
	}
	res += sumEvenGrandparentHelp(root.Left, root.Val%2, p)
	res += sumEvenGrandparentHelp(root.Right, root.Val%2, p)
	return res
}
