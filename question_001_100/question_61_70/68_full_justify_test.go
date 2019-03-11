package question_61_70

import (
	"reflect"
	"strings"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test#1",
			args{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16},
			[]string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			"test#2",
			args{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16},
			[]string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			"test#3",
			args{[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20},
			[]string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() \n[%v], \n\n[%v]", strings.Join(got, "]\n["), strings.Join(tt.want, "]\n["))
			}
		})
	}
}
