package question_31_40

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test#0", nil, nil},
		{"test#1", []int{1}, []int{1}},
		{"test#2", []int{1, 2}, []int{2, 1}},
		{"test#3", []int{1, 2, 3}, []int{1, 3, 2}},
		{"test#4", []int{3, 2, 1}, []int{1, 2, 3}},
		{"test#5", []int{1, 1, 5}, []int{1, 5, 1}},
		{"test#6", []int{2, 1}, []int{1, 2}},
		{"test#7", []int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{"test#8", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"test#9", []int{1, 3, 2}, []int{2, 1, 3}},
		{"test#10", []int{2, 3, 1}, []int{3, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
