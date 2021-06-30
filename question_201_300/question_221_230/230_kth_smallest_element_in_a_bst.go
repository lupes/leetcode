package question_221_230

// 230. 二叉搜索树中第K小的元素
// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
// Topics: 树 二分查找

import . "github.com/lupes/leetcode/common"

func kthSmallest(root *TreeNode, k int) int {
	_, n := kthSmallestHelper(root, k, 0)
	return n
}

func kthSmallestHelper(root *TreeNode, k, count int) (int, int) {
	if root == nil {
		return count, -1
	}
	leftCount, n := kthSmallestHelper(root.Left, k, count)
	if n != -1 {
		return 0, n
	}
	if leftCount == k-1 {
		return 0, root.Val
	}
	RightCount, n := kthSmallestHelper(root.Right, k, leftCount+1)
	if n != -1 {
		return 0, n
	}
	return RightCount, -1
}
