package question_1191_1200

import (
	"reflect"
	"testing"
)

func Test_minimumAbsDifference(t *testing.T) {
	tests := []struct {
		arr  []int
		want [][]int
	}{
		{[]int{4, 2, 1, 3}, [][]int{{1, 2}, {2, 3}, {3, 4}}},
		{[]int{1, 3, 6, 10, 15}, [][]int{{1, 3}}},
		{[]int{3, 8, -10, 23, 19, -4, -14, 27}, [][]int{{-14, -10}, {19, 23}, {23, 27}}},
		{[]int{40, 11, 26, 27, -20}, [][]int{{26, 27}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minimumAbsDifference(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumAbsDifference() = %#v, want %v", got, tt.want)
			}
		})
	}
}
