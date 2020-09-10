func groupAnagrams(strs []string) [][]string {
	words := make(map[string][]string)
	var result [][]string
	for _, s := range strs {
		count := make([]int, 26)
		for _, c := range s {
			count[c-'a']++
		}
		sb := strings.Builder{}
		for _, k := range count {
			sb.WriteString(fmt.Sprintf("%s#", strconv.Itoa(k)))
		}
		words[sb.String()] = append(words[sb.String()], s)
	}
	for _, v := range words {
		result = append(result, v)
	}
	return result
}
