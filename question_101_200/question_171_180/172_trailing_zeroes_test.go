package question_171_180

import "testing"

func Test_trailingZeroes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test", 1, 0},
		{"test", 5, 1},
		{"test", 10, 2},
		{"test", 15, 3},
		{"test", 20, 4},
		{"test", 25, 6},
		{"test", 125, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
