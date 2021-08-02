package question_201_210

import (
	"testing"
)

func Test_canFinish(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
		{5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canFinish(tt.numCourses, tt.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
