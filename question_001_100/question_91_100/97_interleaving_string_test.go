package question_91_100

import "testing"

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test_1", args{"aabcc", "dbbca", "aadbbcbcac"}, true},
		{"test_2", args{"aabcc", "dbbca", "aadbbbaccc"}, false},
		{"test_3", args{"a", "b", "a"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
