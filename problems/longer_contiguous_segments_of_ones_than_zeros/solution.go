func checkZeroOnes(s string) bool {
	maxones, maxo, sumones, sumo := 0, 0, 0, 0
	numbers := strings.Split(s, "")

	for i := 0; i < len(numbers); i++ {
		if i != 0 {
			if numbers[i] != numbers[i-1] {
				if numbers[i] == "0" {
					sumo = 0
				} else {
					sumones = 0
				}
			}
		}
		if numbers[i] == "1" {
			sumones++
			if sumones > maxones {
				maxones = sumones
			}
		} else {
			sumo++
			if sumo > maxo {
				maxo = sumo
			}
		}

	}
	return maxones > maxo

}