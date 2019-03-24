package question_0011_0020

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		wants []int
		want  int
	}{
		{"test#0", []int{}, []int{}, 0},
		{"test#1", []int{1}, []int{1}, 1},
		{"test#2", []int{1, 1}, []int{1}, 1},
		{"test#3", []int{1, 1, 2}, []int{1, 2}, 2},
		{"test#4", []int{1, 1, 2, 2}, []int{1, 2}, 2},
		{"test#5", []int{1, 1, 2, 2, 3}, []int{1, 2, 3}, 3},
		{"test#6", []int{1, 2}, []int{1, 2}, 2},
		{"test#7", []int{1, 2, 3}, []int{1, 2, 3}, 3},
		{"test#8", []int{1, 1, 2, 2, 3, 4}, []int{1, 2, 3, 4}, 4},
		{"test#9", []int{1, 1, 2, 2, 3, 4}, []int{1, 2, 3, 4}, 4},
		{"test#10", []int{1, 1, 2, 2, 3, 4, 5, 5}, []int{1, 2, 3, 4, 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.nums[0:got], tt.wants) {
				t.Errorf("removeDuplicates() = %v, want %v", tt.nums[0:got], tt.wants)
			}
		})
	}
}
