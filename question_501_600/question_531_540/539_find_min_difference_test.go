package question_531_540

import "testing"

func Test_findMinDifference(t *testing.T) {
	tests := []struct {
		timePoints []string
		want       int
	}{
		{[]string{"23:59", "00:00"}, 1},
		{[]string{"23:59", "23:50", "00:00"}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findMinDifference(tt.timePoints); got != tt.want {
				t.Errorf("findMinDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
