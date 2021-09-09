package question_61_70

import (
	"reflect"
	"strings"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	tests := []struct {
		words    []string
		maxWidth int
		want     []string
	}{
		{
			[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16,
			[]string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16,
			[]string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20,
			[]string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := fullJustify(tt.words, tt.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() \n[%v], \n\n[%v]", strings.Join(got, "]\n["), strings.Join(tt.want, "]\n["))
			}
		})
	}
}
