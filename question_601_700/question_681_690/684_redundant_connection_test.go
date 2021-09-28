package question_681_690

import (
	"reflect"
	"testing"
)

func Test_findRedundantConnection(t *testing.T) {
	tests := []struct {
		edges [][]int
		want  []int
	}{
		{[][]int{{1, 2}, {1, 3}, {2, 3}}, []int{2, 3}},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, []int{1, 4}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findRedundantConnection(tt.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
