func removeVowels(S string) string {
	var x []rune
	for _, i2 := range S {
		switch i2 {
		case 'a':
		case 'e':
		case 'i':
		case 'o':
		case 'u':
		continue
		default:
			x = append(x,i2)
		}
	}
	return string(x)
}