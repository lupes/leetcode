package question_41_50

import (
	"testing"
)

func Test_firstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test#0", []int{}, 1},
		{"test#1", []int{1, 2, 0}, 3},
		{"test#2", []int{3, 4, -1, 1}, 2},
		{"test#3", []int{7, 8, 9, 11, 12}, 1},
		{"test#4", []int{0, 2, 2, 1, 1}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
