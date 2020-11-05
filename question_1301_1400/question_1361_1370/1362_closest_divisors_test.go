package question_1361_1370

import (
	"reflect"
	"testing"
)

func Test_closestDivisors(t *testing.T) {
	tests := []struct {
		num  int
		want []int
	}{
		{8, []int{3, 3}},
		{123, []int{5, 25}},
		{999, []int{25, 40}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := closestDivisors(tt.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
