package question_621_630

import (
	"testing"
)

func Test_scheduleCourse(t *testing.T) {
	tests := []struct {
		courses [][]int
		want    int
	}{
		{[][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}, 3},
		{[][]int{{1, 2}}, 1},
		{[][]int{{3, 2}, {4, 3}}, 0},
		{[][]int{{100, 200}, {200, 1300}, {100, 1300}, {1000, 1250}, {2000, 3200}}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := scheduleCourse(tt.courses); got != tt.want {
				t.Errorf("scheduleCourse() = %v, want %v", got, tt.want)
			}
		})
	}
}
