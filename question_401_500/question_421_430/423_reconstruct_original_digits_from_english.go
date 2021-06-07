package question_421_430

import (
	"bytes"
)

// 423. 从英文中重建数字
// https://leetcode-cn.com/problems/reconstruct-original-digits-from-english
// Topics: 数学

func originalDigits(s string) string {
	// zero one     two three four five six seven eight nine
	// z    o-0-2-4 w   h-8   u    f-4  x   s-6   g     i-5-6-8
	var flag, res, nums = make([]int, 'z'+1), bytes.Buffer{}, make([]int, 10)
	for _, c := range s {
		flag[c]++
	}
	nums[0] = flag['z']
	nums[2] = flag['w']
	nums[4] = flag['u']
	nums[1] = flag['o'] - nums[0] - nums[2] - nums[4]
	nums[5] = flag['f'] - nums[4]
	nums[6] = flag['x']
	nums[7] = flag['s'] - nums[6]
	nums[8] = flag['g']
	nums[3] = flag['h'] - nums[8]
	nums[9] = flag['i'] - nums[5] - nums[6] - nums[8]

	for n, num := range nums {
		for i := 0; i < num; i++ {
			res.WriteByte(byte(n) + '0')
		}
	}

	return res.String()
}
