package question_841_850

import (
	"testing"
)

func Test_canVisitAllRooms(t *testing.T) {
	tests := []struct {
		rooms [][]int
		want  bool
	}{
		{[][]int{{1}, {2}, {3}, {}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canVisitAllRooms(tt.rooms); got != tt.want {
				t.Errorf("canVisitAllRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
