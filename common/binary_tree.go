package common

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (self *TreeNode) String() string {
	if self.Left != nil && self.Right != nil {
		return fmt.Sprintf("&TreeNode{Val:%d, Left:%+v, Right:%+v},", self.Val, self.Left, self.Right)
	} else if self.Left != nil {
		return fmt.Sprintf("&TreeNode{Val:%d, Left:%+v},", self.Val, self.Left)
	} else if self.Right != nil {
		return fmt.Sprintf("&TreeNode{Val:%d, Right:%+v},", self.Val, self.Right)
	} else {
		return fmt.Sprintf("&TreeNode{Val:%d}", self.Val)
	}
}
