package question_671_680

// 671. 二叉树中第二小的节点
// https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree
// Topics: 树

import (
	"sort"

	. "github.com/lupes/leetcode/common"
)

func findSecondMinimumValue(root *TreeNode) int {
	var flag = make(map[int]struct{})
	findSecondMinimumValueHelper(root, flag)
	var arr = make([]int, 0, len(flag))
	for n, _ := range flag {
		arr = append(arr, n)
	}
	sort.Ints(arr)
	if len(arr) > 1 {
		return arr[1]
	}
	return -1
}

func findSecondMinimumValueHelper(root *TreeNode, flag map[int]struct{}) {
	if root == nil {
		return
	}
	flag[root.Val] = struct{}{}
	findSecondMinimumValueHelper(root.Left, flag)
	findSecondMinimumValueHelper(root.Right, flag)
}
