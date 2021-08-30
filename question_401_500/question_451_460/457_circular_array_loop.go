package question_451_460

// 457. 环形数组循环
// https://leetcode-cn.com/problems/circular-array-loop
// Topics: 数组 双指针

func circularArrayLoop(nums []int) bool {
	l, cur, next := len(nums), 0, 0
	for start, cnt := 0, 0; start < l; {
		cur = next
		next = (cur + (nums[cur] % l) + l) % l
		cnt += 1
		if cnt > l || nums[cur]*nums[next] < 0 || next == cur {
			start++
			next = start
			cnt = 0
		} else if next == start {
			return true
		}
	}
	return false
}
