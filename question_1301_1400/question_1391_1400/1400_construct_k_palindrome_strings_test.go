package question_1371_1380

import (
	"testing"
)

func Test_canConstruct(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want bool
	}{
		{"annabelle", 1, true},
		{"annabelle", 2, true},
		{"annabelle", 3, true}, // anna elle b
		{"annabelle", 4, true}, // aa nn elle b
		{"annabelle", 5, true}, // aa nn ee ll b
		{"annabelle", 6, true}, // a a nn ee ll b
		{"annabelle", 7, true}, // a a n n ee ll b
		{"annabelle", 8, true}, // a a n n e e ll b
		{"annabelle", 9, true}, // a a n n e e l l b
		{"leetcode", 3, false},
		{"leetcode", 6, true},        // l t c o d eee
		{"yzyzyzyzyzyzyzy", 2, true}, // l t c o d eee
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canConstruct(tt.s, tt.k); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
