package question_281_290

import (
	"reflect"
	"testing"
)

func Test_addOperators(t *testing.T) {
	tests := []struct {
		num    string
		target int
		want   []string
	}{
		{"123", 6, []string{"1+2+3", "1*2*3"}},
		//{"232", 8, []string{"2+3*2", "2*3+2"}},
		//{"105", 5, []string{"10-5", "1*0+5"}},
		{"00", 0, []string{"0+0", "0-0", "0*0"}},
		//{"3456237490", 9191, nil},
		//{"0123456789", 9191, nil},
		{"123456789", 45, nil},
		{"2147483648", -2147483648, nil},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := addOperators(tt.num, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOperators() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_operator(t *testing.T) {
	tests := []struct {
		num  string
		want int
	}{
		{"12-3*4*5-6-78-9", 45},
		{"1+2+3", 6},
		{"2147483648", 45},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := operator(tt.num); got != tt.want {
				t.Errorf("operator() = %v, want %v", got, tt.want)
			}
		})
	}
}
