package question_11_20

func isValid(s string) bool {
	var stack []uint8
	var size int
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, uint8(c))
			size++
		} else if size > 0 &&
			((c == ')' && stack[size-1] == '(') ||
				(c == ']' && stack[size-1] == '[') ||
				(c == '}' && stack[size-1] == '{')) {
			stack = stack[:size-1]
			size--
		} else if c == ')' || c == ']' || c == '}' {
			return false
		}
	}
	return size == 0
}
