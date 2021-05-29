func checkOnesSegment(s string) bool {
    sumones:= 0
	numbers := strings.Split(s, "")

	for i := 0; i < len(numbers); i++ {
        if i != 0 &&  numbers[i] != numbers[i-1]  && numbers[i] == "1" && sumones > 0 {
					return false
		}
        
		if numbers[i] == "1" {
			sumones++
		} 
	}
	return true
}