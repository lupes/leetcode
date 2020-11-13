package question_1331_1340

// 1359. 有效的快递序列数目
// https://leetcode-cn.com/problems/count-all-valid-pickup-and-delivery-options/
// Topics: 数学 动态规划

func countOrders(n int) int {
	var flag = make([]int, n+1)
	flag[1] = 1
	for i := 2; i <= n; i++ {
		t := (i-1)*2 + 1
		flag[i] = (t * (t + 1) / 2 * flag[i-1]) % 1000000007
	}
	return flag[n]
}
