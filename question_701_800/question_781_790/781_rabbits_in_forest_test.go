package question_781_790

import (
	"testing"
)

func Test_numRabbits(t *testing.T) {
	tests := []struct {
		answers []int
		want    int
	}{
		{[]int{1, 1, 2}, 5},
		{[]int{10, 10, 10}, 11},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numRabbits(tt.answers); got != tt.want {
				t.Errorf("numRabbits() = %v, want %v", got, tt.want)
			}
		})
	}
}
