package question_11_20

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test#1", args{"()"}, true},
		{"test#2", args{"()[]{}"}, true},
		{"test#3", args{"(]"}, false},
		{"test#4", args{"([)]"}, false},
		{"test#5", args{"{[]}"}, true},
		{"test#6", args{"}"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
