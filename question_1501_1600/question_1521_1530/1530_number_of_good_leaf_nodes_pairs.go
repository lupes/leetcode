package question_1521_1530

// 1530. 好叶子节点对的数量
// https://leetcode-cn.com/problems/number-of-good-leaf-nodes-pairs/
// Topics: 树 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func countPairs(root *TreeNode, distance int) int {
	if distance == 1 {
		return 0
	}
	_, res := countPairsDFS(root, distance)
	return res
}

func countPairsDFS(root *TreeNode, distance int) ([]int, int) {
	var resFlag = make([]int, distance+1)
	if root.Left == nil && root.Right == nil {
		resFlag[1] = 1
		return resFlag, 0
	}
	var leftFlag, rightFlag []int
	var leftRes, rightRes, res int
	if root.Left != nil {
		leftFlag, leftRes = countPairsDFS(root.Left, distance)
	}
	if root.Right != nil {
		rightFlag, rightRes = countPairsDFS(root.Right, distance)
	}

	res += leftRes + rightRes

	if leftFlag != nil && rightFlag != nil {
		for n, flag := range leftFlag {
			for i := 1; i <= distance-n; i++ {
				res += flag * rightFlag[i]
			}
		}
	}

	for i := 1; i < distance && leftFlag != nil; i++ {
		resFlag[i+1] += leftFlag[i]
	}

	for i := 1; i < distance && rightFlag != nil; i++ {
		resFlag[i+1] += rightFlag[i]
	}

	return resFlag, res
}
