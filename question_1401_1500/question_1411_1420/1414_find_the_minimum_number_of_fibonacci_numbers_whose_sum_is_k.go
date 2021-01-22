package question_1411_1420

// 1414. 和为 K 的最少斐波那契数字数目
// https://leetcode-cn.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/
// Topics: 栈 设计

func findMinFibonacciNumbers(k int) int {
	var fibonacci = []int{1, 1}
	for i := 2; fibonacci[i-1] < k; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}

	var res = 0
	for j := len(fibonacci) - 1; j >= 0 && k > 0; j-- {
		if fibonacci[j] <= k {
			k, res = k-fibonacci[j], res+1
		}
	}

	return res
}
