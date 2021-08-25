package question_791_800

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	tests := []struct {
		graph [][]int
		want  [][]int
	}{
		{
			[][]int{{1, 2}, {3}, {3}, {}},
			[][]int{{0, 1, 3}, {0, 2, 3}},
		},
		{
			[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
			[][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}},
		},
		{
			[][]int{{1}, {}},
			[][]int{{0, 1}},
		},
		{
			[][]int{{1, 2, 3}, {2}, {3}, {}},
			[][]int{{0, 1, 2, 3}, {0, 2, 3}, {0, 3}},
		},
		{
			[][]int{{1, 3}, {2}, {3}, {}},
			[][]int{{0, 1, 2, 3}, {0, 3}},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := allPathsSourceTarget(tt.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %#v, want %v", got, tt.want)
			}
		})
	}
}
