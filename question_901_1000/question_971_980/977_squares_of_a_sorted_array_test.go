package question_971_980

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	tests := []struct {
		A    []int
		want []int
	}{
		{[]int{-2, 0, 1}, []int{0, 1, 4}},
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sortedSquares(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
