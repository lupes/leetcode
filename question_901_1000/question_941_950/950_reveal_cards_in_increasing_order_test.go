package question_941_950

import (
	"reflect"
	"testing"
)

func Test_deckRevealedIncreasing(t *testing.T) {
	tests := []struct {
		deck []int
		want []int
	}{
		{[]int{17, 13, 11, 2, 3, 5, 7}, []int{2, 13, 3, 11, 5, 17, 7}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := deckRevealedIncreasing(tt.deck); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deckRevealedIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
