package question_1301_1310

import (
	"reflect"
	"testing"
)

func Test_xorQueries(t *testing.T) {
	tests := []struct {
		arr     []int
		queries [][]int
		want    []int
	}{
		{[]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}, []int{2, 7, 14, 8}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := xorQueries(tt.arr, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("xorQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
