package question_121_130

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"test", nil, 0},
		{"test", []int{1}, 0},
		{"test", []int{1, 2}, 1},
		{"test", []int{2, 1}, 0},
		{"test", []int{2, 1, 2}, 1},
		{"test", []int{7, 1, 5, 3, 6, 4}, 5},
		{"test", []int{7, 6, 4, 3, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
