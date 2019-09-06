package question_221_230

import "testing"

func Test_calculate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test", "1 + 1", 2},
		{"test", " 2 - 1 + 2", 3},
		{"test", "2-1+2", 3},
		{"test", "(1+(4+5+2)-3)+(6+8)", 23},
		{"test", "2-(5-6)", 3},
		{"test", "2-(5+7)", -10},
		{"test", "2-(5+7)+12", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
