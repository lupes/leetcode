package question_421_430

// 427. 建立四叉树
// https://leetcode-cn.com/problems/construct-quad-tree
// Topics:

type Node2 struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node2
	TopRight    *Node2
	BottomLeft  *Node2
	BottomRight *Node2
}

func construct(grid [][]int) *Node2 {
	return nil
}
