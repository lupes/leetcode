package question_141_150

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) String() string {
	return fmt.Sprintf("V: %d, L: %p, R: %p", tn.Val, tn.Left, tn.Right)
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stacks := []*TreeNode{root}
	for len(stacks) > 0 {
		size := len(stacks) - 1
		node := stacks[size]
		stacks = stacks[:size]
		if node.Right != nil {
			stacks = append(stacks, node.Right)
		}

		if node.Left != nil {
			stacks = append(stacks, node.Left)
		}

		res = append(res, node.Val)
		//res = append([]int{node.Val}, res...)
	}
	fmt.Printf("%v\n", res)
	return res
}
