package question_41_50

import "testing"

func Test_myPow(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		n    int
		want float64
	}{
		{"test#0", 2, 0, 1},
		{"test#1", 2, 10, 1024},
		{"test#2", 2.1, 3, 9.261},
		{"test#3", 2, -2, 0.25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.x, tt.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
