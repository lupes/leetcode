package question_261_270

import "testing"

func Test_nthUglyNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test", 1, 1},
		{"test", 2, 2},
		{"test", 3, 3},
		{"test", 4, 4},
		{"test", 5, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.n); got != tt.want {
				t.Errorf("nthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
