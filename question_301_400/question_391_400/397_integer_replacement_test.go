package question_391_400

import "testing"

func Test_integerReplacement(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test", 1, 0},
		{"test", 2, 1},
		{"test", 3, 2},
		{"test", 8, 3},
		{"test", 7, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerReplacement(tt.n); got != tt.want {
				t.Errorf("integerReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
