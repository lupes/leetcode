package question_231_240

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"test", 1, true},
		{"test", 2, true},
		{"test", 3, false},
		{"test", 4, true},
		{"test", 5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
