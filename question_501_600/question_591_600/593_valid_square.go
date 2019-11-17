package question_591_600

import (
	"sort"
)

// 593. 有效的正方形
// https://leetcode-cn.com/problems/valid-square/
// Topics: 数学

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	var a []int
	a = append(a, distance(p1, p2))
	a = append(a, distance(p1, p3))
	a = append(a, distance(p1, p4))
	a = append(a, distance(p2, p3))
	a = append(a, distance(p2, p4))
	a = append(a, distance(p3, p4))
	sort.Ints(a)
	if a[0] != 0 && a[0] == a[3] && a[0]*2 == a[4] && a[4] == a[5] {
		return true
	}
	return false
}

func distance(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}
