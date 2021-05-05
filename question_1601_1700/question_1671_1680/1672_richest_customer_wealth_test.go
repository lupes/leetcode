package question_1671_1680

import (
	"testing"
)

func Test_maximumWealth(t *testing.T) {
	tests := []struct {
		accounts [][]int
		want     int
	}{
		{[][]int{{1, 2, 3}, {3, 2, 1}}, 6},
		{[][]int{{1, 5}, {7, 3}, {3, 5}}, 10},
		{[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, 17},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maximumWealth(tt.accounts); got != tt.want {
				t.Errorf("maximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
