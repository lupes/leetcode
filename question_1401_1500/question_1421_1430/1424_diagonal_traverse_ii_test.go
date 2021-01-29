package question_1421_1430

import (
	"reflect"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	tests := []struct {
		nums [][]int
		want []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 4, 2, 7, 5, 3, 8, 6, 9}},
		{[][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}, []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16}},
		{[][]int{{1, 2, 3}, {4}, {5, 6, 7}, {8}, {9, 10, 11}}, []int{1, 4, 2, 5, 3, 8, 6, 9, 7, 10, 11}},
		{[][]int{{1, 2, 3, 4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findDiagonalOrder(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
