package mian_shi

// 面试题 10.01. 合并排序的数组
// https://leetcode-cn.com/problems/sorted-merge-lcci/
// Topics: 数组 双指针 排序

func merge(A []int, m int, B []int, n int) {
	for ai, bi, end := m-1, n-1, m+n-1; bi >= 0; {
		if ai >= 0 && A[ai] >= B[bi] {
			A[end] = A[ai]
			ai--
		} else {
			A[end] = B[bi]
			bi--
		}
		end--
	}
}
