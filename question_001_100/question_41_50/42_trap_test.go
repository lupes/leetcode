package question_41_50

import "testing"

func Test_trap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"test#0", nil, 0},
		{"test#1", []int{1, 0}, 0},
		{"test#2", []int{1, 0, 1}, 1},
		{"test#3", []int{1, 0, 2}, 1},
		{"test#4", []int{1, 0, 1, 2}, 1},
		{"test#5", []int{2, 0, 1, 2}, 3},
		{"test#6", []int{2, 0, 0, 2}, 4},
		{"test#7", []int{3, 0, 0, 2}, 4},
		{"test#8", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
