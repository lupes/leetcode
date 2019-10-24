package question_521_530

// 530. 二叉搜索树的最小绝对差
// https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	_, _, r := getMinimumDifferenceDfs(root)
	return r
}

func getMinimumDifferenceDfs(root *TreeNode) (int, int, int) {
	var l, r, res1, res2 = root.Val, root.Val, -1, -1

	if root.Left != nil {
		l1, r1, t := getMinimumDifferenceDfs(root.Left)
		res1 = root.Val - r1
		if t != -1 && t < res1 {
			res1 = t
		}
		l = l1
	}

	if root.Right != nil {
		l2, r2, t := getMinimumDifferenceDfs(root.Right)
		res2 = l2 - root.Val
		if t != -1 && t < res2 {
			res2 = t
		}
		r = r2
	}
	if res1 == -1 && res2 == -1 {
		return l, r, -1
	} else if res1 == -1 && res2 != -1 {
		return l, r, res2
	} else if res1 != -1 && res2 == -1 {
		return l, r, res1
	} else if res1 > res2 {
		return l, r, res2
	} else {
		return l, r, res1
	}
}
