package question_881_890

// 888. 公平的糖果交换
// https://leetcode-cn.com/problems/fair-candy-swap
// Topics: 数组

func fairCandySwap(A []int, B []int) []int {
	var bMap, aSum, bSum = make(map[int]struct{}, len(B)), 0, 0
	for _, a := range A {
		aSum += a
	}
	for _, b := range B {
		bMap[b] = struct{}{}
		bSum += b
	}
	d := (aSum - bSum) / 2
	for _, a := range A {
		if _, ok := bMap[a-d]; ok {
			return []int{a, a - d}
		}
	}
	return nil
}
