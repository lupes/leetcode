package question_731_740

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	tests := []struct {
		T    []int
		want []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := dailyTemperatures(tt.T); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
