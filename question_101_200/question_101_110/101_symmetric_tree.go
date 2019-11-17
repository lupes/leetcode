package question_101_110

// 101. 对称二叉树
// https://leetcode-cn.com/problems/symmetric-tree/
// Topics: 树 深度优先搜索 广度优先搜索

func isSymmetric(root *TreeNode) bool {
	var next = []*TreeNode{root}
	for len(next) != 0 {
		var son []*TreeNode
		for _, t := range next {
			if t != nil {
				son = append(son, t.Left, t.Right)
			}
		}
		if len(son)%2 != 0 {
			return false
		}
		for i, l := 0, len(son); i < l/2; i++ {
			if (son[i] != nil && son[l-i-1] == nil) ||
				(son[i] == nil && son[l-i-1] != nil) ||
				(son[i] != nil && son[l-i-1] != nil && son[i].Val != son[l-i-1].Val) {
				return false
			}
		}
		next = son
	}
	return true
}
