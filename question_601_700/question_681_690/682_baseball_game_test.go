package question_681_690

import (
	"testing"
)

func Test_calPoints(t *testing.T) {
	tests := []struct {
		ops  []string
		want int
	}{
		{[]string{"5", "2", "C", "D", "+"}, 30},
		{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
		{[]string{"1"}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := calPoints(tt.ops); got != tt.want {
				t.Errorf("calPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
