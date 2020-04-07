package question_951_960

import (
	"reflect"
	"testing"
)

func Test_prisonAfterNDays(t *testing.T) {
	tests := []struct {
		cells []int
		N     int
		want  []int
	}{
		{[]int{0, 1, 0, 1, 1, 0, 0, 1}, 7, []int{0, 0, 1, 1, 0, 0, 0, 0}},
		{[]int{0, 1, 0, 1, 1, 0, 0, 1}, 70, []int{0, 0, 0, 0, 1, 1, 0, 0}},
		{[]int{0, 1, 0, 1, 1, 0, 0, 1}, 69, []int{0, 0, 1, 1, 0, 1, 0, 0}},
		{[]int{0, 1, 0, 1, 1, 0, 0, 1}, 690, []int{0, 0, 0, 0, 0, 1, 0, 0}},
		{[]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000, []int{0, 0, 1, 1, 1, 1, 1, 0}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := prisonAfterNDays(tt.cells, tt.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prisonAfterNDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
