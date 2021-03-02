package question_1491_1500

import (
	"testing"
)

func Test_average(t *testing.T) {
	tests := []struct {
		salary []int
		want   float64
	}{
		{[]int{1000, 2000, 3000}, 2000},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := average(tt.salary); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
