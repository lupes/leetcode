package question_491_500

import "testing"

func Test_findMaximizedCapital(t *testing.T) {
	tests := []struct {
		k       int
		W       int
		Profits []int
		Capital []int
		want    int
	}{
		{2, 0, []int{1, 2, 3}, []int{0, 1, 1}, 4},
		{1, 2, []int{1, 2, 3}, []int{1, 1, 2}, 5},
		{11, 11, []int{1, 2, 3}, []int{11, 12, 13}, 17},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findMaximizedCapital(tt.k, tt.W, tt.Profits, tt.Capital); got != tt.want {
				t.Errorf("findMaximizedCapital() = %v, want %v", got, tt.want)
			}
		})
	}
}
