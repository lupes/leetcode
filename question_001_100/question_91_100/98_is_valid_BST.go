package question_91_100

// 98. 验证二叉搜索树
// https://leetcode-cn.com/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	res := isValidBSTDfs(root)
	for i := 1; i < len(res); i++ {
		if res[i-1] >= res[i] {
			return false
		}
	}
	return true
}

func isValidBSTDfs(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	if root.Left != nil {
		res = append(res, isValidBSTDfs(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, isValidBSTDfs(root.Right)...)
	}
	return res
}
