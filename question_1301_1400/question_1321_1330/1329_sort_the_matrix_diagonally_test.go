package question_1321_1330

import (
	"reflect"
	"testing"
)

func Test_diagonalSort(t *testing.T) {
	tests := []struct {
		mat  [][]int
		want [][]int
	}{
		{[][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}, [][]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 3, 3}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := diagonalSort(tt.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diagonalSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
