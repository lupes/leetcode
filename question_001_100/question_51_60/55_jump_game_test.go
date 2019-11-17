package question_51_60

import "testing"

func Test_canJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"test#0", []int{0}, true},
		{"test#1", []int{1}, true},
		{"test#2", []int{2, 3, 1, 1, 4}, true},
		{"test#3", []int{3, 2, 1, 0, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
