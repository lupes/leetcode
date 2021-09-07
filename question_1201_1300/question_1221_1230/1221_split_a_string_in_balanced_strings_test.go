package question_1221_1230

import (
	"testing"
)

func Test_balancedStringSplit(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"RLRRLLRLRL", 4},
		{"RLLLLRRRLR", 3},
		{"LLLLRRRR", 1},
		{"RLRRRLLRLL", 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := balancedStringSplit(tt.s); got != tt.want {
				t.Errorf("balancedStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
