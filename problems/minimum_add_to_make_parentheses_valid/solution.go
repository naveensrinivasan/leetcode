func minAddToMakeValid(S string) int {

	result := 0
	if len(S) == 0 {
		return 0
	}
	makeValidParentheses := func(stack []string, character string, result *int) []string {

		if len(stack) == 0 {
			*result++
			return stack
		}
		if stack[len(stack)-1] != character {
			*result++
		}
		stack = stack[:len(stack)-1]
		return stack
	}
	var stack []string
	for _, char := range S {
		switch string(char) {
		case "{", "(", "[":
			stack = append(stack, string(char))
		case "]":
			stack = makeValidParentheses(stack, "[", &result)
		case "}":
			stack = makeValidParentheses(stack, "{", &result)
		case ")":
			stack = makeValidParentheses(stack, "(", &result)
		}

	}

	return result + len(stack)
}
