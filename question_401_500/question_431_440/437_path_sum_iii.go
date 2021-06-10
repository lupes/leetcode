package question_431_440

// 437. 路径总和 III
// https://leetcode-cn.com/problems/path-sum-iii
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func pathSum(root *TreeNode, sum int) int {
	var flag = make(map[int]int, 0)
	flag[0] = 1
	return pathSumHelper(root, sum, 0, flag)
}

func pathSumHelper(root *TreeNode, sum, prefix int, flag map[int]int) int {
	if root == nil {
		return 0
	}
	prefix += root.Val
	var res = flag[prefix-sum]
	flag[prefix] += 1
	res += pathSumHelper(root.Left, sum, prefix, flag)
	res += pathSumHelper(root.Right, sum, prefix, flag)
	flag[prefix] -= 1
	return res
}
