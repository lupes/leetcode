package question_1301_1310

import (
	"reflect"
	"testing"
)

func Test_sumZero(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{1, []int{0}},
		{2, []int{1, -1}},
		{3, []int{1, -1, 0}},
		{4, []int{1, -1, 2, -2}},
		{5, []int{1, -1, 2, -2, 0}},
		{6, []int{1, -1, 2, -2, 3, -3}},
		{7, []int{1, -1, 2, -2, 3, -3, 0}},
		{8, []int{1, -1, 2, -2, 3, -3, 4, -4}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sumZero(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
