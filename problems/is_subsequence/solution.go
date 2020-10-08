func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	sp, tp := 0, 0
	for tp < len(t) {
		if t[tp] == s[sp] {
			sp++
		}
		if sp == len(s) {
			return true
		}
		tp++
	}
	return false
}
