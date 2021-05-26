func slowestKey(releaseTimes []int, keysPressed string) byte {
	max, char := 0, 0
	for i := 0; i < len(keysPressed); i++ {
		diff := 0
		if i == 0 {
			max = releaseTimes[0]
			char = int(keysPressed[i])
			continue
		} else {
			diff = releaseTimes[i] - releaseTimes[i-1]
		}

		if diff > max {
			char = int(keysPressed[i])
			max = diff
		} else if diff == max && int(keysPressed[i]) > char {
			char = int(keysPressed[i])
		}
	}

	return byte(char)
}