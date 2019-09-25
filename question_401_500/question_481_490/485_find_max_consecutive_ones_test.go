package question_481_490

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test", []int{0}, 0},
		{"test", []int{}, 0},
		{"test", []int{1}, 1},
		{"test", []int{1, 0}, 1},
		{"test", []int{0, 1}, 1},
		{"test", []int{1, 1, 0, 1, 1, 1}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
