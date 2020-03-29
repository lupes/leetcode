package question_911_920

import (
	"testing"
)

func Test_partitionDisjoint(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{[]int{5, 0, 3, 8, 6}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := partitionDisjoint(tt.A); got != tt.want {
				t.Errorf("partitionDisjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
