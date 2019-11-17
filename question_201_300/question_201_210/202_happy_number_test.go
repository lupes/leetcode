package question_201_210

import "testing"

func Test_isHappy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"test", 1, true},
		{"test", 2, false},
		{"test", 3, false},
		{"test", 100, true},
		{"test", 1001, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
