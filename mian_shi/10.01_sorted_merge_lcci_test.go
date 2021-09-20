package mian_shi

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		A []int
		m int
		B []int
		n int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
		{[]int{-1, 0, 0, 3, 3, 3, 0, 0, 0}, 6, []int{1, 2, 2}, 3},
		{[]int{0}, 0, []int{1}, 1},
		{[]int{1}, 1, []int{0}, 0},
		{[]int{2, 0}, 1, []int{1}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			merge(tt.A, tt.m, tt.B, tt.n)
			fmt.Printf("%v\n", tt.A)
		})
	}
}
