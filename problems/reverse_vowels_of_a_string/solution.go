func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	m := make(map[byte]rune)
	m['a'] = 'a'
	m['e'] = 'e'
	m['i'] = 'i'
	m['o'] = 'o'
	m['u'] = 'u'
	m['A'] = 'A'
	m['E'] = 'E'
	m['I'] = 'I'
	m['O'] = 'O'
	m['U'] = 'U'
	result := make([]byte, len(s))
	for left <= right {
		if _, ok := m[s[left]]; !ok {
			result[left] = s[left]
			left++
			continue
		}
		if _, ok := m[s[right]]; !ok {
			result[right] = s[right]
			right--
			continue
		}
		result[left], result[right] = s[right], s[left]
		left++
		right--

	}
	return string(result)
}
