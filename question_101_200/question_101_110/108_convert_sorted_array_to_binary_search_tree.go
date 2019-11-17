package question_101_110

// 108. 将有序数组转换为二叉搜索树
// https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
// Topics: 树 深度优先搜索

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	left := sortedArrayToBST(nums[:l/2])
	right := sortedArrayToBST(nums[l/2+1:])
	return &TreeNode{Val: nums[l/2], Left: left, Right: right}
}
