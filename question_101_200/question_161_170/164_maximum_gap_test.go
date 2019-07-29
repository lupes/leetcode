package question_161_170

import "testing"

func Test_maximumGap(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test", nil, 0},
		{"test", []int{1}, 0},
		{"test", []int{1, 3, 8, 6}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGap(tt.nums); got != tt.want {
				t.Errorf("maximumGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
