func checkInclusion(s1 string, s2 string) bool {
	var x []int
	l := len(s1)
	for i := 0; i<=len(s2) -l;i++  {
		if isAnagram(s2[i:i+l],s1){
			x = append(x,i)
		}
	}
	return len(x) > 0
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	x := make([]int,26)
	for i := 0; i < len(s); i++ {
		x[s[i]-97]++
		x[t[i]-97]--
	}
	for _, i2 := range x {
		if i2!=0{
			return false
		}
	}
	return true
}