package mian_shi

import (
	"reflect"
	"testing"
)

func Test_smallestK(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want []int
	}{
		{[]int{1, 3, 5, 7, 2, 4, 6, 8}, 4, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := smallestK(tt.arr, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestK() = %v, want %v", got, tt.want)
			}
		})
	}
}
