package question_851_860

// 858. 镜面反射
// https://leetcode-cn.com/problems/mirror-reflection
// Topics: 数学

func mirrorReflection(p int, q int) int {
	g := gcd(p, q)
	p, q = p/g, q/g
	p, q = p%2, q%2
	if p == 1 && q == 1 {
		return 1
	} else if p == 1 {
		return 0
	} else {
		return 2
	}
}

func gcd(p, q int) int {
	if q == 0 {
		return p
	}
	return gcd(q, p%q)
}
