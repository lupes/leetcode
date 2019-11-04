package question_591_600

import (
	"reflect"
	"testing"
)

func Test_findRestaurant(t *testing.T) {
	tests := []struct {
		list1 []string
		list2 []string
		want  []string
	}{
		{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}, []string{"Shogun"}},
		{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}, []string{"Shogun"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findRestaurant(tt.list1, tt.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRestaurant() = %v, want %v", got, tt.want)
			}
		})
	}
}
