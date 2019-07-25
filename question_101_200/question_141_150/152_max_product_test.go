package question_141_150

import "testing"

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test", nil, 0},
		{"test", []int{1}, 1},
		{"test", []int{2, 3, -2, 4}, 6},
		{"test", []int{-2, 0, 1, 0, -1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
