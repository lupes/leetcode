package question_1251_1260

import (
	"reflect"
	"testing"
)

func Test_shiftGrid(t *testing.T) {
	tests := []struct {
		grid [][]int
		k    int
		want [][]int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1,
			[][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			[][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4,
			[][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9,
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			[][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}}, 23,
			[][]int{{6}, {5}, {1}, {2}, {3}, {4}, {7}},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := shiftGrid(tt.grid, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shiftGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
