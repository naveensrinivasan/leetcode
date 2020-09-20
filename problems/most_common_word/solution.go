func mostCommonWord(paragraph string, banned []string) string {
	bannedTable := make(map[string]string, len(banned))
	wordCountTable := make(map[string]int)
	regualrExpression := regexp.MustCompile("[^a-zA-Z ]")
	para := regualrExpression.ReplaceAllString(strings.ToLower(paragraph), " ")
	greatest := 0
	commonWord := ""

	for _, word := range banned {
		bannedTable[word] = word
	}

	for _, word := range strings.Split(para, " ") {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}
		if _, ok := bannedTable[word]; !ok {
			wordCountTable[word] += 1
			if wordCountTable[word] > greatest {
				greatest = wordCountTable[word]
				commonWord = word
			}
		}
	}

	return commonWord
}
