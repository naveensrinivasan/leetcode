func wordPattern(pattern string, str string) bool {
	m := make(map[int32]string)
	words := strings.Split(str, " ")
    if len(pattern) != len(words) {
		return false
	}
	for i, i2 := range pattern {
		if value, ok := m[i2]; ok {
			if value != words[i] {
				return false
			}
			continue
		}
		for _, i4 := range m {
			if i4 == words[i] {
				return false
			}
		}
		m[i2] = words[i]
	}
	return true
}