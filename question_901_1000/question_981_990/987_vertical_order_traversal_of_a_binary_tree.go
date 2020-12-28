package question_981_990

// 987. 二叉树的垂序遍历
// https://leetcode-cn.com/problems/vertical-order-traversal-of-a-binary-tree
// Topics: 树 哈希表

import (
	"sort"

	. "github.com/lupes/leetcode/common"
)

func verticalTraversal(root *TreeNode) [][]int {
	flag := make(map[int][]int, 0)
	verticalTraversalHelper(root, flag, 0, 1)

	var keys []int
	for k, _ := range flag {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var res = make([][]int, len(keys))
	for i, k := range keys {
		sort.Ints(flag[k])
		res[i] = make([]int, 0, len(flag[k]))
		for _, n := range flag[k] {
			res[i] = append(res[i], n%10000)
		}
	}
	return res
}

func verticalTraversalHelper(root *TreeNode, flag map[int][]int, n, y int) {
	if root == nil {
		return
	}
	flag[n] = append(flag[n], root.Val+y*10000)
	verticalTraversalHelper(root.Left, flag, n-1, y+1)
	verticalTraversalHelper(root.Right, flag, n+1, y+1)
}
