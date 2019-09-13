package question_411_420

// 412. Fizz Buzz
// https://leetcode-cn.com/problems/fizz-buzz/

import "strconv"

func fizzBuzz(n int) []string {
	var res, s = make([]string, n), ""
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			s = "FizzBuzz"
		} else if i%5 == 0 {
			s = "Buzz"
		} else if i%3 == 0 {
			s = "Fizz"
		} else {
			s = strconv.Itoa(i)
		}
		res[i-1] = s
	}
	return res
}
