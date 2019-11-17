package question_101_110

// 103. 二叉树的锯齿形层次遍历
// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var next = []*TreeNode{root}
	for i := 1; len(next) > 0; i++ {
		var nextNext []*TreeNode
		var t []int
		for _, node := range next {
			if i%2 == 0 {
				t = append([]int{node.Val}, t...)
			} else {
				t = append(t, node.Val)
			}

			if node.Left != nil {
				nextNext = append(nextNext, node.Left)
			}
			if node.Right != nil {
				nextNext = append(nextNext, node.Right)
			}
		}
		next = nextNext
		res = append(res, t)
	}

	return res
}
