package question_1201_1210

import (
	"testing"
)

func Test_uniqueOccurrences(t *testing.T) {

	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := uniqueOccurrences(tt.arr); got != tt.want {
				t.Errorf("uniqueOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
