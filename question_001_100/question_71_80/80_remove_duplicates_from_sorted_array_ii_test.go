package question_71_80

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test#0", []int{}, []int{}},
		{"test#1", []int{1}, []int{1}},
		{"test#2", []int{1, 1}, []int{1, 1}},
		{"test#3", []int{1, 1, 1}, []int{1, 1}},
		{"test#4", []int{1, 1, 1, 2}, []int{1, 1, 2}},
		{"test#5", []int{1, 1, 1, 2, 2}, []int{1, 1, 2, 2}},
		{"test#6", []int{1, 1, 1, 2, 2, 3}, []int{1, 1, 2, 2, 3}},
		{"test#7", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, []int{0, 0, 1, 1, 2, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); !reflect.DeepEqual(tt.nums[:got], tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", tt.nums[:got], tt.want)
			}
		})
	}
}
