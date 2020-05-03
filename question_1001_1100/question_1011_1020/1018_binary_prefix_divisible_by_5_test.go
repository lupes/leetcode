package question_1011_1020

import (
	"reflect"
	"testing"
)

func Test_prefixesDivBy5(t *testing.T) {
	tests := []struct {
		A    []int
		want []bool
	}{
		{[]int{0, 1, 1}, []bool{true, false, false}},
		{[]int{1, 1, 1, 0, 1}, []bool{false, false, false, false, false}},
		{[]int{0, 1, 1, 1, 1, 1}, []bool{true, false, false, false, true, false}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := prefixesDivBy5(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixesDivBy5() = %v, want %v", got, tt.want)
			}
		})
	}
}
