package question_851_860

import (
	"reflect"
	"testing"
)

func Test_loudAndRich(t *testing.T) {
	tests := []struct {
		richer [][]int
		quiet  []int
		want   []int
	}{
		{[][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}, []int{5, 5, 2, 5, 4, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := loudAndRich(tt.richer, tt.quiet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loudAndRich() = %v, want %v", got, tt.want)
			}
		})
	}
}
