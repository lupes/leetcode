package question_2021_2030

import (
	"reflect"
	"testing"
)

func Test_construct2DArray(t *testing.T) {
	tests := []struct {
		original []int
		m        int
		n        int
		want     [][]int
	}{
		{[]int{1, 2, 3, 4}, 2, 2, [][]int{{1, 2}, {3, 4}}},
		{[]int{1, 2, 3}, 1, 3, [][]int{{1, 2, 3}}},
		{[]int{1, 2}, 1, 1, nil},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := construct2DArray(tt.original, tt.m, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("construct2DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
