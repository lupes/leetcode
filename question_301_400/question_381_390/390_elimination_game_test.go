package question_381_390

import "testing"

func Test_lastRemaining(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test", 1, 1},
		{"test", 2, 2},
		{"test", 3, 2},
		{"test", 4, 2},
		{"test", 5, 2},
		{"test", 6, 4},
		{"test", 9, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastRemaining(tt.n); got != tt.want {
				t.Errorf("lastRemaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
