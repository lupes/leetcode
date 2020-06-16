package question_1231_1240

import (
	"testing"
)

func Test_checkStraightLine(t *testing.T) {
	tests := []struct {
		coordinates [][]int
		want        bool
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}}, true},
		{[][]int{{1, 1}, {2, 2}, {3, 4}}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := checkStraightLine(tt.coordinates); got != tt.want {
				t.Errorf("checkStraightLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
