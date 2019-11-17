package question_91_100

import . "github.com/lupes/leetcode/common"

// 95. 不同的二叉搜索树 II
// https://leetcode-cn.com/problems/unique-binary-search-trees-ii/

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTreesHelper(1, n)
}

func generateTreesHelper(start, end int) []*TreeNode {
	var res []*TreeNode
	if end < start {
		return append(res, nil)
	}
	for i := start; i <= end; i++ {
		for _, leftTree := range generateTreesHelper(start, i-1) {
			for _, rightTree := range generateTreesHelper(i+1, end) {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  leftTree,
					Right: rightTree,
				})
			}
		}
	}
	return res
}
