package question_691_700

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		words []string
		k     int
		want  []string
	}{
		{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2, []string{"i", "love"}},
		{[]string{"i", "love", "leetcode", "i", "love", "coding", "coding"}, 2, []string{"coding", "i"}},
		{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4, []string{"the", "is", "sunny", "day"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := topKFrequent(tt.words, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
