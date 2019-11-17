package question_491_500

// 508. 出现次数最多的子树元素和
// https://leetcode-cn.com/problems/most-frequent-subtree-sum/
// Topics: 树 哈希表

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var tmp = make(map[int]int)
	var res = make(map[int][]int)
	var max = 0
	findFrequentTreeSumDfs(root, tmp)

	for k, v := range tmp {
		if v > max {
			max = v
		}
		res[v] = append(res[v], k)
	}
	return res[max]
}

func findFrequentTreeSumDfs(root *TreeNode, res map[int]int) (sum int) {
	sum = root.Val
	if root.Left != nil {
		sum += findFrequentTreeSumDfs(root.Left, res)
	}
	if root.Right != nil {
		sum += findFrequentTreeSumDfs(root.Right, res)
	}
	res[sum]++
	return sum
}
