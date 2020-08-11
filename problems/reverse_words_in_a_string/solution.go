func reverseWords(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)/2; i++ {
		words[i], words[(len(words)-1)-(i*1)] = words[(len(words)-1)-(i*1)], words[i]
	}
	return strings.Join(words, " ")
}
