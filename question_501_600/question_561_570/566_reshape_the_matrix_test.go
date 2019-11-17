package question_561_570

import (
	"reflect"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	tests := []struct {
		nums [][]int
		r    int
		c    int
		want [][]int
	}{
		{[][]int{{1, 2}, {3, 4}}, 1, 4, [][]int{{1, 2, 3, 4}}},
		{[][]int{{1, 2}, {3, 4}}, 2, 4, [][]int{{1, 2}, {3, 4}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := matrixReshape(tt.nums, tt.r, tt.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixReshape() = %v, want %v", got, tt.want)
			}
		})
	}
}
