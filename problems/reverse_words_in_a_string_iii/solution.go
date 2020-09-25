func reverseWords(s string) string {
	words := splitWords(s)
	result := ""
	for _, w := range words {
		w = trimSpace(w)
		b := []byte(w)
		revWord(b)
		result += " " + string(b)
	}
	return result[1:]
}
func splitWords(s string) []string {
	start, end := 0, 0
	result := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			end++
			continue
		}
		end++
		result = append(result, s[start:end])
		start = i + 1
	}
	result = append(result, s[start:end])
	return result
}
func revWord(s []byte) {
	l := len(s)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}
func trimSpace(s string) string {
	result := ""
	for _, v := range s {
		if v != 32 {
			result += string(v)
		}
	}
	revWord([]byte(result))
	return result
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-1-i], s[i]
	}
}
