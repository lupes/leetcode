package question_781_790

import (
	"reflect"
	"testing"
)

func Test_kthSmallestPrimeFraction(t *testing.T) {
	tests := []struct {
		A    []int
		K    int
		want []int
	}{
		{[]int{1, 2, 3, 5}, 3, []int{2, 5}},
		{[]int{1, 7}, 1, []int{1, 7}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := kthSmallestPrimeFraction(tt.A, tt.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthSmallestPrimeFraction() = %v, want %v", got, tt.want)
			}
		})
	}
}
