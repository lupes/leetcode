package question_491_500

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		S    int
		want int
	}{
		{"test", []int{1, 1, 1, 1, 1}, 3, 5},
		{"test", []int{1, 1, 1, 1}, 3, 0},
		{"test", []int{1, 1, 1}, 3, 1},
		{"test", []int{1, 1, 3}, 3, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.nums, tt.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
