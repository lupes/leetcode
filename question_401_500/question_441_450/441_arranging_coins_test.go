package question_441_450

import "testing"

func Test_arrangeCoins(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test", 0, 0},
		{"test", 1, 1},
		{"test", 2, 1},
		{"test", 3, 2},
		{"test", 4, 2},
		{"test", 5, 2},
		{"test", 6, 3},
		{"test", 7, 3},
		{"test", 10, 4},
		{"test", 11, 4},
		{"test", 15, 5},
		{"test", 16, 5},
		{"test", 22, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrangeCoins(tt.n); got != tt.want {
				t.Errorf("arrangeCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
