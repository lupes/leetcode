package question_1021_1030

import (
	"testing"
)

func Test_twoCitySchedCost(t *testing.T) {
	tests := []struct {
		costs [][]int
		want  int
	}{
		{[][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}, 110},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := twoCitySchedCost(tt.costs); got != tt.want {
				t.Errorf("twoCitySchedCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
