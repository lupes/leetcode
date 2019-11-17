package question_51_60

import "testing"

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test#0", []int{0}, 0},
		{"test#1", []int{1, 1}, 2},
		{"test#2", []int{1, 1, -1}, 2},
		{"test#3", []int{1, 1, -1, 3}, 4},
		{"test#4", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
