package question_871_880

// 872. 叶子相似的树
// https://leetcode-cn.com/problems/leaf-similar-trees
// Topics: 树 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	t1 := leafSimilarHelper(root1)
	t2 := leafSimilarHelper(root2)
	if len(t1) != len(t2) {
		return false
	}

	for i, n := range t1 {
		if n != t2[i] {
			return false
		}
	}
	return true
}

func leafSimilarHelper(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var res = leafSimilarHelper(root.Left)
	return append(res, leafSimilarHelper(root.Right)...)
}
