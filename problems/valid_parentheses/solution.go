func isValid(s string) bool {
   var sl []string
	for _, i2 := range s {
		switch i2 {
		case '{':
			sl = append(sl, "{")
		case '}':
			if len(sl) == 0 || sl[len(sl)-1] != "{" {
				return false
			}
			sl = sl[:(len(sl) - 1)]
		case '[':
			sl = append(sl, "[")
		case ']':
			if len(sl) == 0 || sl[len(sl)-1] != "[" {
				return false
			}
			sl = sl[:(len(sl) - 1)]
		case '(':
			sl = append(sl, "(")
		case ')':
			if len(sl) == 0 || sl[len(sl)-1] != "(" {
				return false
			}
			sl = sl[:(len(sl) - 1)]
		}
	}
	return len(sl) == 0
}