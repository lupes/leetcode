package question_161_170

import (
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test", nil, -1},
		{"test", []int{0}, 0},
		{"test", []int{1, 0}, 0},
		{"test", []int{0, 1}, 1},
		{"test", []int{0, 1, 0}, 1},
		{"test", []int{1, 2, 3, 1}, 2},
		{"test", []int{1, 2, 1, 3, 5, 6, 4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
