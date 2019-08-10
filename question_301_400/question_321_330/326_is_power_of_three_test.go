package question_321_330

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"test", 0, false},
		{"test", 1, false},
		{"test", 3, true},
		{"test", 27, true},
		{"test", 45, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
