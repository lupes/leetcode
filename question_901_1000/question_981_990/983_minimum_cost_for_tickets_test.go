package question_981_990

import (
	"testing"
)

func Test_mincostTickets(t *testing.T) {
	tests := []struct {
		days  []int
		costs []int
		want  int
	}{
		{[]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}, 17},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := mincostTickets(tt.days, tt.costs); got != tt.want {
				t.Errorf("mincostTickets() = %v, want %v", got, tt.want)
			}
		})
	}
}
