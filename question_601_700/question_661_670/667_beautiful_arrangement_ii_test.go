package question_661_670

import (
	"reflect"
	"testing"
)

func Test_constructArray(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want []int
	}{
		{3, 1, []int{1, 2, 3}},
		{3, 2, []int{1, 3, 2}},
		{4, 3, []int{1, 4, 2, 3}},
		{10, 3, []int{1, 4, 2, 3, 5, 6, 7, 8, 9, 10}},
		{10, 4, []int{1, 5, 2, 4, 3, 6, 7, 8, 9, 10}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := constructArray(tt.n, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructArray() = %#v, want %v", got, tt.want)
			}
		})
	}
}
