func isValid(s string) bool {
	box := []string{}
	for _, v := range s {
		switch string(v) {
		case "{", "[", "(":
			box = append(box, string(v))
		case "}":
			if len(box) == 0 || box[len(box)-1] != "{" {
				return false
			}
			box = box[:len(box)-1]
		case "]":
			if len(box) == 0 || box[len(box)-1] != "[" {
				return false
			}
			box = box[:len(box)-1]
		case ")":
			if len(box) == 0 || box[len(box)-1] != "(" {
				return false
			}
			box = box[:len(box)-1]
		}

	}
	return len(box) == 0
}
