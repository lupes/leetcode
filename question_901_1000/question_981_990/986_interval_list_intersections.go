package question_981_990

// 986. 区间列表的交集
// https://leetcode-cn.com/problems/interval-list-intersections
// Topics: 双指针

func intervalIntersection(A [][]int, B [][]int) [][]int {
	a, b, res := 0, 0, make([][]int, 0)
	for a < len(A) && b < len(B) {
		aStart, aEnd, bStart, bEnd := A[a][0], A[a][1], B[b][0], B[b][1]
		if aStart > bEnd {
			b++
		} else if bStart > aEnd {
			a++
		} else if bStart >= aStart && aEnd >= bEnd {
			res = append(res, []int{bStart, bEnd})
			b++
		} else if aStart >= bStart && bEnd >= aEnd {
			res = append(res, []int{aStart, aEnd})
			a++
		} else if aStart >= bStart && aEnd >= bEnd {
			res = append(res, []int{aStart, bEnd})
			b++
		} else if bStart >= aStart && bEnd >= aEnd {
			res = append(res, []int{bStart, aEnd})
			a++
		}
	}
	return res
}
