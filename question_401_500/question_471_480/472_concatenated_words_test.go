package question_471_480

import (
	"reflect"
	"testing"
)

func Test_findAllConcatenatedWordsInADict(t *testing.T) {
	tests := []struct {
		words []string
		want  []string
	}{
		{
			[]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"},
			[]string{"dogcatsdog", "catsdogcats", "ratcatdogcat"},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findAllConcatenatedWordsInADict(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllConcatenatedWordsInADict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addWord(t *testing.T) {
	var root = &node{
		end:   false,
		child: make(map[int32]*node),
	}
	tests := []string{"cat", "cats", "dog"}
	for _, tt := range tests {
		addWord(root, tt)
	}
}
