package question_111_120

// 113. 路径总和 II
// https://leetcode-cn.com/problems/path-sum-ii/
// Topics: 树 深度优先搜索

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	if root.Left == nil && root.Right == nil && sum == root.Val {
		return [][]int{{root.Val}}
	}
	leftRes := pathSum(root.Left, sum-root.Val)
	for _, r := range leftRes {
		res = append(res, append([]int{root.Val}, r...))
	}
	rightRes := pathSum(root.Right, sum-root.Val)
	for _, r := range rightRes {
		res = append(res, append([]int{root.Val}, r...))
	}
	return res
}
