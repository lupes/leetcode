package question_1041_1050

import (
	"testing"
)

func Test_lastStoneWeightII(t *testing.T) {
	tests := []struct {
		stones []int
		want   int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := lastStoneWeightII(tt.stones); got != tt.want {
				t.Errorf("lastStoneWeightII() = %v, want %v", got, tt.want)
			}
		})
	}
}
