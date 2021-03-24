package question_1551_1560

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findSmallestSetOfVertices(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
		want  []int
	}{
		{6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}, []int{0, 3}},
		{5, [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}, []int{0, 2, 3}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := findSmallestSetOfVertices(tt.n, tt.edges)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSmallestSetOfVertices() = %v, want %v", got, tt.want)
			}
		})
	}
}
