package question_31_40

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{nil, nil},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{2, 3, 1}, []int{3, 1, 2}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			nextPermutation(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
