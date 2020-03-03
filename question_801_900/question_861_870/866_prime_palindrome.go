package question_861_870

import (
	"math"
)

// 866. 回文素数
// https://leetcode-cn.com/problems/prime-palindrome
// Topics: 数学

func primePalindrome(N int) int {
	check := []int{2, 2, 2, 3, 5, 5, 7, 7, 11, 11, 11, 11}
	if N <= 11 {
		return check[N]
	}
	if N > 9989899 {
		return 100030001
	}
	if N%2 == 0 {
		N++
	}
	var t1, t2 = 0, 0
	for i := N; ; i += 2 {
		t1, t2 = i, 0
		for t1 > 0 {
			t2 = t2*10 + t1%10
			t1 /= 10
		}
		if i != t2 {
			goto Next
		}
		for j := 3; j <= int(math.Sqrt(float64(i))); j += 2 {
			if i%j == 0 {
				goto Next
			}
		}
		return i
	Next:
	}
}
