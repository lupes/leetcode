package question_41_50

import "testing"

func Test_jump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test#0", []int{0}, 0},
		{"test#1", []int{1}, 0},
		{"test#2", []int{2, 1}, 1},
		{"test#3", []int{2, 1, 3}, 1},
		{"test#4", []int{2, 1, 3, 1}, 2},
		{"test#5", []int{2, 1, 3, 1, 1}, 2},
		{"test#6", []int{2, 1, 3, 1, 1, 1}, 2},
		{"test#7", []int{2, 3, 1, 1, 4}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.nums); got != tt.want {
				t.Errorf("jump(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
