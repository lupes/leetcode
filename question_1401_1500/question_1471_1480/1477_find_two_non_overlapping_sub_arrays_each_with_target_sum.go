package question_1471_1480

// 1477. 找两个和为目标值且不重叠的子数组
// https://leetcode-cn.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/
// Topics: 动态规划

const max = 1000000

func minSumOfLengths(arr []int, target int) int {
	var l = len(arr)

	var mapLeft, flagLeft, sumLeft = make(map[int]int), make([]int, l+1), 0
	flagLeft[0] = max
	for i := 0; i < l; i++ {
		sumLeft += arr[i]
		t := max
		if j, ok := mapLeft[sumLeft-target]; ok && sumLeft-target > 0 {
			t = i - j
		} else if sumLeft-target == 0 {
			t = i + 1
		}
		if t < flagLeft[i] {
			flagLeft[i+1] = t
		} else {
			flagLeft[i+1] = flagLeft[i]
		}
		mapLeft[sumLeft] = i
	}
	flagLeft = flagLeft[1:]

	var mapRight, flagRight, sumRight = make(map[int]int), make([]int, l+1), 0
	flagRight[l] = max
	for i := l - 1; i >= 0; i-- {
		sumRight += arr[i]
		t := max
		if j, ok := mapRight[sumRight-target]; ok && sumRight-target > 0 {
			t = j - i
		} else if sumRight-target == 0 {
			t = l - i
		}
		if t < flagRight[i+1] {
			flagRight[i] = t
		} else {
			flagRight[i] = flagRight[i+1]
		}
		mapRight[sumRight] = i
	}
	flagRight = flagRight[:l]

	var min = max
	for i := 0; i < l-1; i++ {
		if flagLeft[i]+flagRight[i+1] < min {
			min = flagLeft[i] + flagRight[i+1]
		}
	}

	if min == max {
		return -1
	}

	return min
}
