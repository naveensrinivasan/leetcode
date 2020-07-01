
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

func findAnagrams(s string, p string) []int {
	var x []int
	l := len(p)
	for i := 0; i <= len(s)-l; i++ {
		if isAnagram(s[i:i+l],p){
			x = append(x,i)
		}
	}
	return x
}
