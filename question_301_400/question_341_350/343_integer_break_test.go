package question_341_350

import "testing"

func Test_integerBreak(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test", 2, 1},
		{"test", 3, 2},
		{"test", 4, 4},
		{"test", 7, 12},
		{"test", 10, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerBreak(tt.n); got != tt.want {
				t.Errorf("integerBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
