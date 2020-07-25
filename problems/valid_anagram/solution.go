func isAnagram(s string, t string) bool {
	anagram := func(a, b string) bool {
		if len(a) != len(b) {
			return false
		}
		m := make([]int, 26)
		for i := range a {
			m[a[i]-97]++
			m[b[i]-97]--
		}
		for _, i2 := range m {
			if i2 != 0 {
				return false
			}
		}
		return true
	}
	return anagram(s, t)
}
