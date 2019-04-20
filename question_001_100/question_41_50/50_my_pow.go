package question_41_50

func myPow(x float64, n int) float64 {
	if x == 1 || n == 0 {
		return 1
	}
	if x == -1 {
		if n%2 == 0 {
			return 1
		} else {
			return -1
		}
	}
	var flag bool
	var t = n
	if n < 0 {
		flag = true
		t = -n
	}
	var res = 1.0
	for i := 0; i < t; i++ {
		res *= x
	}
	if flag {
		res = 1 / res
	}
	return res
}
