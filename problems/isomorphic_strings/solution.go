func isIsomorphic(s string, t string) bool {
    m := make(map[int32]int32)
	for i, i2 := range s {
		if value, ok := m[i2]; ok {
			if value != rune(t[i]) {
				return false
			}
			continue
		}
		for _, i4 := range m {
			if i4 == rune(t[i]) {
				return false
			}
		}
		m[i2] = rune(t[i])
	}
	return true
}