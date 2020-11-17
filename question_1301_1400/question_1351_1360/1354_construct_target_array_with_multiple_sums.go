package question_1331_1340

import (
	"sort"
)

// 1354. 多次求和构造目标数组
// https://leetcode-cn.com/problems/construct-target-array-with-multiple-sums/
// Topics: 贪心算法

func isPossible(target []int) bool {
	l := len(target) - 1
	if l == 0 {
		return target[0] == 1
	}
	sort.Ints(target)

	sum := 0
	for _, n := range target {
		sum += n
	}
	for {
		if sum-target[l] == 1 {
			return true
		}
		t := target[l]*2 - sum
		if target[l]/(sum-target[l]) > 2 {
			t = target[l] % (sum - target[l])
		}

		if t <= 0 {
			return false
		}
		sum -= target[l] - t
		target[l] = t
		isPossibleHelper(target)
		if target[l] == 1 {
			return true
		}
	}
}

func isPossibleHelper(target []int) {
	for i := len(target) - 1; i > 0; i-- {
		if target[i-1] > target[i] {
			target[i-1], target[i] = target[i], target[i-1]
		}
	}
}
