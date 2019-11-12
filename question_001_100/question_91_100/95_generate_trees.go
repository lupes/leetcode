package question_91_100

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
		leftTrees := generateTreesHelper(start, i-1)
		rightTrees := generateTreesHelper(i+1, end)
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
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
