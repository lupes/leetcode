package question_1301_1310

// 1305. 两棵二叉搜索树中的所有元素
// https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees/
// Topics: 排序 树

import (
	"sort"

	. "github.com/lupes/leetcode/common"
)

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var res []int
	middle(root1, &res)
	middle(root2, &res)
	sort.Ints(res)
	return res
}

func middle(root *TreeNode, res *[]int) {
	if root != nil {
		middle(root.Left, res)
		*res = append(*res, root.Val)
		middle(root.Right, res)
	}
}
