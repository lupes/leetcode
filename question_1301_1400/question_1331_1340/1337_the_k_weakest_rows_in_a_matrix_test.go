package question_1311_1320

import (
	"reflect"
	"testing"
)

func Test_kWeakestRows(t *testing.T) {
	tests := []struct {
		mat  [][]int
		k    int
		want []int
	}{
		{[][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}, 3, []int{2, 0, 3}},
		{[][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 2, []int{0, 2}},
		{[][]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}}, 2, []int{0, 1}},
		{[][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 1, 1, 1, 1}}, 3, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := kWeakestRows(tt.mat, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kWeakestRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
