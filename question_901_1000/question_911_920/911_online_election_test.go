package question_911_920

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		persons []int
		times   []int
		want    TopVotedCandidate
	}{
		{[]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30}, TopVotedCandidate{}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := Constructor(tt.persons, tt.times); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopVotedCandidate_Q(t *testing.T) {
	tests := []struct {
		persons []int
		times   []int
		q       []int
		want    []int
	}{
		{
			[]int{0, 1, 1, 0, 0, 1, 0},
			[]int{0, 5, 10, 15, 20, 25, 30},
			[]int{3, 12, 25, 15, 24, 8},
			[]int{0, 1, 1, 0, 0, 1},
		},
		{
			[]int{0, 1, 0, 1, 1},
			[]int{24, 29, 31, 76, 81},
			[]int{28, 24, 29, 77, 30, 25, 76, 75, 81, 80},
			[]int{0, 0, 1, 1, 1, 0, 1, 0, 1, 1},
		},
	}
	for _, tt := range tests {
		this := Constructor(tt.persons, tt.times)
		fmt.Printf("%v\n", this.ts)

		t.Run("test", func(t *testing.T) {
			var got []int
			for _, q := range tt.q {
				got = append(got, this.Q(q))
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nargs = %+v\ngot = %+v\n want %+v", tt.q, got, tt.want)
			}
		})
	}
}
