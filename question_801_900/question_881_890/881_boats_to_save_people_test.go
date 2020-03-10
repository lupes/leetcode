package question_881_890

import (
	"testing"
)

func Test_numRescueBoats(t *testing.T) {
	tests := []struct {
		people []int
		limit  int
		want   int
	}{
		{[]int{1, 2}, 3, 1},
		{[]int{1, 2, 2, 3}, 3, 3},
		{[]int{3, 5, 3, 4}, 5, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numRescueBoats(tt.people, tt.limit); got != tt.want {
				t.Errorf("numRescueBoats() = %v, want %v", got, tt.want)
			}
		})
	}
}
