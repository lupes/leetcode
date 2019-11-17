package question_111_120

import "testing"

func Test_minimumTotal(t *testing.T) {
	triangle := [][]int{
		{2},
		{3, 4}, // 3 = 0-1, 0 4 = 1-1, 1
		{6, 5, 7},
		{4, 1, 8, 3},
		{8, 9, 1, 1, 5},
	}
	tests := []struct {
		name     string
		triangle [][]int
		want     int
	}{
		{"test", triangle, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
