func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	chars := strings.Split(pattern, "")

	if len(words) != len(chars) {
		return false
	}

	h := make(map[string]string)

	for i, c := range chars {
		if value, ok := h[c]; !ok {
			for _, v := range h {
				if v == words[i] {
					return false
				}
			}
			h[c] = words[i]
			continue
		} else {
			if value != words[i] {
				return false
			}
		}
	}

	return true

}
