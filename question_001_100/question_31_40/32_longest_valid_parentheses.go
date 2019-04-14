package question_31_40

func longestValidParentheses(s string) int {
	size := len(s)
	if size < 2 {
		return 0
	}
	flags := make([]int, size)
	max := 0
	for i := 1; i < size; i++ {
		if s[i] == ')' {
			if i-flags[i-1]-1 >= 0 && s[i-flags[i-1]-1] == '(' {
				flags[i] = flags[i-1] + 2
			}
			if i-flags[i] >= 0 {
				flags[i] += flags[i-flags[i]]
			}
			if flags[i] > max {
				max = flags[i]
			}
		}
	}
	return max
}
