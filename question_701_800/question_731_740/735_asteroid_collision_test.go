package question_731_740

import (
	"reflect"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	tests := []struct {
		asteroids []int
		want      []int
	}{
		{[]int{5, 10, -5}, []int{5, 10}},
		{[]int{8, -8}, []int{}},
		{[]int{10, 2, -5}, []int{10}},
		{[]int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := asteroidCollision(tt.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
