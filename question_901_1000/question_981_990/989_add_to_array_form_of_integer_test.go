package question_981_990

import (
	"reflect"
	"testing"
)

func Test_addToArrayForm(t *testing.T) {
	tests := []struct {
		A    []int
		K    int
		want []int
	}{
		{[]int{1, 2, 0, 0}, 34, []int{1, 2, 3, 4}},
		{[]int{2, 7, 4}, 181, []int{4, 5, 5}},
		{[]int{2, 1, 5}, 806, []int{1, 0, 2, 1}},
		{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[]int{0}, 23, []int{2, 3}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := addToArrayForm(tt.A, tt.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addToArrayForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
