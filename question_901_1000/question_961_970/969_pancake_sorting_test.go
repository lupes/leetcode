package question_961_970

import (
	"reflect"
	"testing"
)

func Test_pancakeSort(t *testing.T) {
	tests := []struct {
		A    []int
		want []int
	}{
		{[]int{3, 2, 4, 1}, []int{3, 4, 2, 3, 1, 2}},
		{[]int{1, 2, 3}, nil},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := pancakeSort(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pancakeSort() = %#v, want %+v", got, tt.want)
			}
		})
	}
}
