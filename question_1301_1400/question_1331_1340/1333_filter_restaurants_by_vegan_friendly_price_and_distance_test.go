package question_1311_1320

import (
	"reflect"
	"testing"
)

func Test_filterRestaurants(t *testing.T) {
	tests := []struct {
		restaurants   [][]int
		veganFriendly int
		maxPrice      int
		maxDistance   int
		want          []int
	}{
		{
			[][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}},
			1, 50, 10, []int{3, 1, 5},
		},
		{
			[][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}},
			0, 50, 10, []int{4, 3, 2, 1, 5},
		},
		{
			[][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}},
			0, 30, 3, []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := filterRestaurants(tt.restaurants, tt.veganFriendly, tt.maxPrice, tt.maxDistance)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterRestaurants() = %v, want %v", got, tt.want)
			}
		})
	}
}
