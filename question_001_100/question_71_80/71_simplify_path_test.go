package question_71_80

import "testing"

func Test_simplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"test0", "/", "/"},
		{"test1", "/./", "/"},
		{"test2", "/../", "/"},
		{"test3", "//", "/"},
		{"test4", "//a/", "/a"},
		{"test5", "//a/./", "/a"},
		{"test6", "//a/../", "/"},
		{"test7", "/home//foo/", "/home/foo"},
		{"test8", "/a/./b/../../c/", "/c"},
		{"test9", "/a/../../b/../c//.//", "/c"},
		{"test10", "/a//b////c/d//././/..", "/a/b/c"},
		{"test11", "/.../", "/..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
