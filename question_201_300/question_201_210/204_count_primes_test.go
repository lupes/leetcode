package question_201_210

import "testing"

func Test_countPrimes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test", 1, 0},
		{"test", 2, 0},
		{"test", 3, 1},
		{"test", 4, 2},
		{"test", 5, 2},
		{"test", 6, 3},
		{"test", 7, 3},
		{"test", 10, 4},
		{"test", 49979, 5130},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
