func checkInclusion(s1 string, s2 string) bool {
	for i := 0; i <= len(s2)-len(s1); i++ {
		if isthereasubstring(s1, s2[i:len(s1)+i]) {
			return true
		}
	}
	return false
}

func isthereasubstring(s1, s2 string) bool {
	var alphabetCounter [26]rune
	for i, i2 := range s1 {
		alphabetCounter[i2-97]++
		alphabetCounter[s2[i]-97]--
	}
	for _, i2 := range alphabetCounter {
		if i2 != 0 {
			return false
		}
	}
	return true
}