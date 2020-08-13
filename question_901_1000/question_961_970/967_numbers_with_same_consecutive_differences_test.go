package question_961_970

import (
	"reflect"
	"testing"
)

func Test_numsSameConsecDiff(t *testing.T) {
	tests := []struct {
		N    int
		K    int
		want []int
	}{
		{3, 7, []int{181, 292, 707, 818, 929}},
		{2, 1, []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98}},
		{2, 0, []int{11, 22, 33, 44, 55, 66, 77, 88, 99}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numsSameConsecDiff(tt.N, tt.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numsSameConsecDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
