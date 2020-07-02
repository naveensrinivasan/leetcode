func isValid(s string) bool {
	var l []string
	for _, v := range s {
		switch v {
		case '{', '(','[':
			l = append(l, string(v))
		case '}', ')', ']':
			if len(l) == 0 ||
				(string(v) == "}" && l[len(l)-1] != "{" ||
					string(v) == "]" && l[len(l)-1] != "[" ||
					string(v) == ")" && l[len(l)-1] != "(") {
				return false
			}
			l = l[:len(l) -1]

		}
	}
	return len(l) == 0
}