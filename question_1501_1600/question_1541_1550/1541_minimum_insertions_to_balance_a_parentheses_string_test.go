package question_1541_1550

import (
	"testing"
)

func Test_minInsertions(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"(()))", 1},
		{"())", 0},
		{"))())(", 3},
		{"((((((", 12},
		{")))))))", 5},
		{"(()))(()))()())))", 4},
		// 0 0 1 1 0 0 1 1 0 1 0 1 1
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minInsertions(tt.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}
		})
	}
}
