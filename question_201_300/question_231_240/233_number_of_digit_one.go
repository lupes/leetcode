package question_231_240

import (
	"fmt"
	"math"
)

// 233. 数字 1 的个数
// https://leetcode-cn.com/problems/number-of-digit-one
// Topics: 数学

var m = []int{0, 1, 19, 271, 3439, 40951, 468559, 468559*9 + 100000}

func countDigitOne(n int) int {
	h := int(math.Log10(float64(n)))
	t := n
	var res = 0
	for i := 0; t > 0; i++ {
		w := int(math.Pow10(h - i))
		d := t / w
		t %= w
		if d > 1 {
			res += (d-1)*m[h-i] + w
		} else if d == 1 {
			res += t + m[h-i] + 1
			break
		}

		fmt.Printf("%d %d %d\n", t/w, t%w, res)
	}
	fmt.Printf("%d\n", res)
	return res
}
