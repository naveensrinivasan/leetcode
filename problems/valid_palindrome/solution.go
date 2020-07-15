func isPalindrome(s string) bool {
    i:= 0
	j:= len(s) -1
	var IsLetter = regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString
	for i<j{
		for i <j && !IsLetter(string(s[i])){
			i++
		}
		for i< j && !IsLetter(string(s[j])){
			j--
		}
		if strings.ToLower(string(s[i])) != strings.ToLower( string(s[j])){
			return false
		}
		i++
		j--
	}
	return true
}