func findAnagrams(s string, p string) []int {
	/*
		O(n)
		Logic for figuring out the anagram
		Map with every char
			- Increment the counter from the original word
			- Decrement the counter from the compared string
			- if they aren't same return false
		Sliding window
		- Loop for the original word from 0 to (len(s) - len(p))
	*/
	anagram := func(a, b string) bool {
		if len(a) != len(b) {
			return false
		}
		m := make([]int, 26)
		for i, _ := range a {
			m[a[i]-97]++
			m[b[i]-97]--
		}
		for _, i2 := range m {
			if i2 != 0 {
				return false
			}
		}
		return true
	}
	var result []int
	if len(s) == 0 || len(p) == 0 {
		return nil
	}
	//Sliding window
	for i := 0; i <= (len(s) - len(p)); i++ {
		if anagram(s[i:i+len(p)], p) {
			result = append(result, i)
		}
	}
	return result
}
