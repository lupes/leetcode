package question_1441_1450

import (
	"reflect"
	"testing"
)

func Test_simplifiedFractions(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{1, nil},
		{2, []string{"1/2"}},
		{3, []string{"1/2", "1/3", "2/3"}},
		{4, []string{"1/2", "1/3", "2/3", "1/4", "3/4"}},
		{5, []string{"1/2", "1/3", "2/3", "1/4", "3/4", "1/5", "2/5", "3/5", "4/5"}},
		{6, []string{"1/2", "1/3", "2/3", "1/4", "3/4", "1/5", "2/5", "3/5", "4/5", "1/6", "5/6"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := simplifiedFractions(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simplifiedFractions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simplifiedFractionsGCD(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{2, 3, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := simplifiedFractionsGCD(tt.a, tt.b); got != tt.want {
				t.Errorf("simplifiedFractionsGCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
