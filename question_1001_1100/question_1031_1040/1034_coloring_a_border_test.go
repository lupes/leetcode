package question_1031_1040

import (
	"reflect"
	"testing"
)

func Test_colorBorder(t *testing.T) {
	tests := []struct {
		grid  [][]int
		r0    int
		c0    int
		color int
		want  [][]int
	}{
		{[][]int{{1, 1}, {1, 2}}, 0, 0, 3, [][]int{{3, 3}, {3, 2}}},
		{[][]int{{1, 2, 2}, {2, 3, 2}}, 0, 1, 3, [][]int{{1, 3, 3}, {2, 3, 3}}},
		{[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}},
		{[][]int{{1, 2, 1}, {1, 2, 2}, {2, 2, 1}}, 1, 1, 2, [][]int{{1, 2, 1}, {1, 2, 2}, {2, 2, 1}}},
		{[][]int{{3, 3, 3, 3, 1, 2}, {3, 2, 1, 1, 1, 3}, {3, 3, 3, 1, 1, 3}, {3, 1, 3, 1, 1, 2}, {3, 3, 3, 2, 3, 1}, {1, 2, 1, 2, 3, 2}}, 4, 1, 1, [][]int{{1, 1, 1, 1, 1, 2}, {1, 2, 1, 1, 1, 3}, {1, 1, 1, 1, 1, 3}, {1, 1, 1, 1, 1, 2}, {1, 1, 1, 2, 3, 1}, {1, 2, 1, 2, 3, 2}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := colorBorder(tt.grid, tt.r0, tt.c0, tt.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("colorBorder() = %#v, want %v", got, tt.want)
			}
		})
	}
}
