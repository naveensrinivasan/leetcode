func maxPower(s string) int {
	max, counter := 0, 0

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			if counter == 0 {
				counter = 2
			} else {
				counter++
			}

		} else {
			if counter > max {
				max = counter
			}
			counter = 0
		}
	}
	if counter > max {
		max = counter
	}
	if max == 0 {
		return 1
	}
	return max
}