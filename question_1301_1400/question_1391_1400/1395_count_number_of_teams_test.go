package question_1371_1380

import (
	"testing"
)

func Test_numTeams(t *testing.T) {
	tests := []struct {
		rating []int
		want   int
	}{
		{[]int{2, 5, 3, 4, 1}, 3},
		{[]int{2, 1, 3}, 0},
		{[]int{1, 2, 3, 4}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numTeams(tt.rating); got != tt.want {
				t.Errorf("numTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
