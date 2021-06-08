package question_421_430

import (
	"fmt"
)

// 427. 建立四叉树
// https://leetcode-cn.com/problems/construct-quad-tree
// Topics:

type Node2 struct {
	IsLeaf      bool
	Val         bool
	TopLeft     *Node2
	TopRight    *Node2
	BottomLeft  *Node2
	BottomRight *Node2
}

func (self Node2) String() string {
	return fmt.Sprintf("[L:%v V:%v TL:%s TR:%s BL:%s BR:%s]", self.IsLeaf, self.Val, self.TopLeft, self.TopRight, self.BottomLeft, self.BottomRight)
}

func construct(grid [][]int) *Node2 {
	return quadTree(grid, 0, len(grid), 0, len(grid))
}

func quadTree(grid [][]int, h, hl, w, wl int) *Node2 {
	if wl-w == 1 {
		return &Node2{
			Val:    grid[h][w] == 1,
			IsLeaf: true,
		}
	}
	root := &Node2{Val: true, IsLeaf: false}

	diff := (wl - w) / 2
	root.TopLeft = quadTree(grid, h, hl-diff, w, wl-diff)
	root.TopRight = quadTree(grid, h, hl-diff, w+diff, wl)
	root.BottomLeft = quadTree(grid, h+diff, hl, w, wl-diff)
	root.BottomRight = quadTree(grid, h+diff, hl, w+diff, wl)

	if root.TopLeft.IsLeaf && root.TopRight.IsLeaf && root.BottomLeft.IsLeaf && root.BottomRight.IsLeaf &&
		((root.TopLeft.Val && root.TopRight.Val && root.BottomLeft.Val && root.BottomRight.Val) ||
			(!root.TopLeft.Val && !root.TopRight.Val && !root.BottomLeft.Val && !root.BottomRight.Val)) {
		root.IsLeaf = true
		root.Val = root.TopLeft.Val
		root.TopLeft = nil
		root.TopRight = nil
		root.BottomLeft = nil
		root.BottomRight = nil
	}

	return root
}
