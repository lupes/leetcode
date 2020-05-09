package question_1021_1030

import (
	"reflect"
	"testing"
)

func Test_allCellsDistOrder(t *testing.T) {
	tests := []struct {
		R    int
		C    int
		r0   int
		c0   int
		want [][]int
	}{
		{1, 2, 0, 0, [][]int{{0, 0}, {0, 1}}},
		{2, 2, 0, 1, [][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}}},
		{2, 3, 1, 2, [][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := allCellsDistOrder(tt.R, tt.C, tt.r0, tt.c0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allCellsDistOrder() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
