package question_1281_1290

import (
	"testing"
)

func Test_findSpecialInteger(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 2, 2, 6, 6, 6, 6, 7, 10}, 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findSpecialInteger(tt.arr); got != tt.want {
				t.Errorf("findSpecialInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
