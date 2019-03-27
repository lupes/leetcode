package question_0011_0020

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test#0", args{nil, nil}, nil},
		{"test#1", args{nil, []string{}}, nil},
		{"test#2", args{"", []string{}}, nil},
		{"test#3", args{"", []string{""}}, nil},
		{"test#4", args{"i", []string{"i"}}, []int{0}},
		{"test#5", args{"ii", []string{"i"}}, []int{0, 1}},
		{"test#6", args{"barfoothefoobarman", []string{"foo", "bar"}}, []int{0, 9}},
		{"test#6", args{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
