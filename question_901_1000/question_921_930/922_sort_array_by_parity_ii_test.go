package question_921_930

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	tests := []struct {
		A    []int
		want []int
	}{
		{[]int{4, 2, 5, 7}, []int{4, 5, 2, 7}},
		{[]int{1, 2, 4, 7}, []int{2, 1, 4, 7}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{2, 1}, []int{2, 1}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sortArrayByParityII(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParityII() = %v, want %v", got, tt.want)
			}
		})
	}
}
