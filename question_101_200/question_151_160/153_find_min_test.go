package question_151_160

import "testing"

func Test_findMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test", nil, 0},
		{"test", []int{0}, 0},
		{"test", []int{0, 1}, 0},
		{"test", []int{1, 0}, 0},
		{"test", []int{3, 4, 5, 1, 2}, 1},
		{"test", []int{4, 5, 6, 7, 0, 1, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
