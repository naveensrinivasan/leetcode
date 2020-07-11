func checkInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	for i := 0; i < (len(s2)-len(s1))+1; i++ {
		if ArePermutations(s1, s2[i:i+len(s1)]) {
			return true
		}
	}
	return false
}



func ArePermutations(s1, s2 string)bool  {
	if len(s1) != len(s2){
		return false
	}
	m := make(map[uint8]int)
	for i:= 0;i<len(s1);i++ {
			m[s1[i]-97]++
			m[s2[i]-97]--
	}
	for _,v := range m{
		if v > 0{
			return false
		}
	}
	return true
}