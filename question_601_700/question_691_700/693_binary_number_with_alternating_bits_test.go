package question_691_700

import "testing"

func Test_hasAlternatingBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"test", 0, true},
		{"test", 1, true},
		{"test", 5, true},
		{"test", 7, false},
		{"test", 10, true},
		{"test", 11, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAlternatingBits(tt.n); got != tt.want {
				t.Errorf("hasAlternatingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
