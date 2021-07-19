package question_821_830

import (
	"testing"
)

func Test_flipgame(t *testing.T) {
	tests := []struct {
		fronts []int
		backs  []int
		want   int
	}{
		{[]int{1, 2, 4, 4, 7}, []int{1, 3, 4, 1, 3}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := flipgame(tt.fronts, tt.backs); got != tt.want {
				t.Errorf("flipgame() = %v, want %v", got, tt.want)
			}
		})
	}
}
