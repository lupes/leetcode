package question_1551_1560

// 1558. 得到目标数组的最少函数调用次数
// https://leetcode-cn.com/problems/minimum-numbers-of-function-calls-to-make-target-array/
// Topics: 贪心算法

func minOperations2(nums []int) int {
	op0, op1 := 0, 0
	for _, n := range nums {
		op0Tmp, op1Tmp := 0, 0
		for n > 0 {
			if n&1 == 1 {
				op0Tmp += 1
				n -= 1
			} else {
				op1Tmp += 1
				n >>= 1
			}
		}
		if op1Tmp > op1 {
			op1 = op1Tmp
		}
		op0 += op0Tmp
	}
	return op0 + op1
}

// 1 3 5 7 9
// 9 => 1:2 2:3
// 7 => 1:3 2:2
// 5
// 0 1 2 3 4 +1
// 0 0 2 2 4 +2
// 0 0 1 1 2 +1
// 0 0 0 0 1 +2

// 3 2 2 4
// 2 2 2 4 1
// 1 1 1 2 1
// 0 0 0 2 3
// 0 0 0 1 1
// 0 0 0 0 1

// 1 5
// 0 4 2
// 0 2 1
// 0 1 1
// 0 0 1
