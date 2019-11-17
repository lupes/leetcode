package question_141_150

import (
	"fmt"
	"reflect"
)

// 149. 直线上最多的点数
// https://leetcode-cn.com/problems/max-points-on-a-line/

func maxPoints(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}

	var res int
	var p1, p2 []int
	var key string

	var r = make(map[string]map[[3]int]struct{})
	for i, p := range points {
		key = fmt.Sprintf("x=%d", p[0])
		if _, ok := r[key]; !ok {
			r[key] = make(map[[3]int]struct{})
		}
		r[key][[3]int{p[0], p[1], i}] = struct{}{}
		key = fmt.Sprintf("y=%d", p[1])
		if _, ok := r[key]; !ok {
			r[key] = make(map[[3]int]struct{})
		}
		r[key][[3]int{p[0], p[1], i}] = struct{}{}
		for j := i + 1; j < len(points); j++ {
			p1, p2 = points[i], points[j]
			if float64(p1[1]-p2[1]) != 0 {
				slope := float64(p1[0]-p2[0]) / float64(p1[1]-p2[1])
				c := float64(p1[1]) - slope*float64(p1[0])
				key := fmt.Sprintf("y=%v*x+%v", slope, c)
				if _, ok := r[key]; !ok {
					r[key] = make(map[[3]int]struct{})
				}
				r[key][[3]int{p1[0], p1[1], i}] = struct{}{}
				r[key][[3]int{p2[0], p2[1], j}] = struct{}{}
			}
			if reflect.DeepEqual(p1, p2) {
				for k, v := range r {
					if _, ok := v[[3]int{p1[0], p1[1], i}]; ok {
						r[k][[3]int{p2[0], p2[1], j}] = struct{}{}
					}
				}
			}
		}
	}
	for _, value := range r {
		if len(value) > res {
			res = len(value)
		}
	}
	return res
}
