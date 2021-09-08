package question_951_960

// 955. 删列造序 II
// https://leetcode-cn.com/problems/delete-columns-to-make-sorted-ii
// Topics: 贪心算法

func minDeletionSize(A []string) int {
	var res, flag, tmp = 0, true, make(map[int]struct{})
	for i := 0; i < len(A[0]) && flag; i++ {
		flag = false
		for j := 1; j < len(A); j++ {
			if A[j][i] == A[j-1][i] {
				flag = true
			} else if A[j][i] < A[j-1][i] {
				k := i - 1
				for ; k >= 0; k-- {
					if _, ok := tmp[k]; ok {
						continue
					} else if A[j][k] > A[j-1][k] {
						break
					}
				}
				if k == -1 {
					res++
					tmp[i] = struct{}{}
					flag = true
					break
				}
			}
		}
	}
	return res
}
