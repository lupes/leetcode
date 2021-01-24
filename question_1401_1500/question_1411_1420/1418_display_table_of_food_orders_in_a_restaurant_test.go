package question_1411_1420

import (
	"reflect"
	"testing"
)

func Test_displayTable(t *testing.T) {
	tests := []struct {
		orders [][]string
		want   [][]string
	}{
		{
			[][]string{{"David", "3", "Ceviche"}, {"Corina", "10", "Beef Burrito"}, {"David", "3", "Fried Chicken"}, {"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"}},
			[][]string{{"Table", "Beef Burrito", "Ceviche", "Fried Chicken", "Water"}, {"3", "0", "2", "1", "0"}, {"5", "0", "1", "0", "1"}, {"10", "1", "0", "0", "0"}},
		},
		{
			[][]string{{"James", "12", "Fried Chicken"}, {"Ratesh", "12", "Fried Chicken"}, {"Amadeus", "12", "Fried Chicken"}, {"Adam", "1", "Canadian Waffles"}, {"Brianna", "1", "Canadian Waffles"}},
			[][]string{{"Table", "Canadian Waffles", "Fried Chicken"}, {"1", "2", "0"}, {"12", "0", "3"}},
		},
		{
			[][]string{{"Laura", "2", "Bean Burrito"}, {"Jhon", "2", "Beef Burrito"}, {"Melissa", "2", "Soda"}},
			[][]string{{"Table", "Bean Burrito", "Beef Burrito", "Soda"}, {"2", "1", "1", "1"}},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := displayTable(tt.orders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("displayTable() = %#v\n want %#v\n", got, tt.want)
			}
		})
	}
}
