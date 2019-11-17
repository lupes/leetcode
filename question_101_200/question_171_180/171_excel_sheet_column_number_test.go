package question_171_180

import "testing"
import "leetcode/question_101_200/question_161_170"

func Test_titleToNumber(t *testing.T) {
	tests := []struct {
		name string
		arg  int
	}{
		{"test", 1},
		{"test", 2},
		{"test", 3},
		{"test", 4},
		{"test", 26},
		{"test", 27},
		{"test", 28},
		{"test", 716},
		{"test", 1716},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := question_161_170.ConvertToTitle(tt.arg)
			if got := titleToNumber(s); got != tt.arg {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.arg)
			}
		})
	}
}
