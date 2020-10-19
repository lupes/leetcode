package question_111_120

// 117. 填充每个节点的下一个右侧节点指针 II
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii
// Topics: 树 深度优先搜索

func connect2(root *Node) *Node {
	next := &Node{Val: 0, Left: nil, Right: nil, Next: nil}
	now, head := next, root
	for head != nil {
		for head != nil {
			if head.Left != nil {
				now.Next, now = head.Left, head.Left
			}
			if head.Right != nil {
				now.Next, now = head.Right, head.Right
			}
			head = head.Next
		}
		head, next.Next, now = next.Next, nil, next
	}
	return root
}
