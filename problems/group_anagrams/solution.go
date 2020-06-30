import (
"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
    x := make(map[string][]string)
	var y [][]string
	for _, i2 := range strs {
		s := sortString(i2)
		if values, ok := x[s]; ok {
			x[s] = append(values, i2)
			continue
		}
		x[s] = []string{i2}
	}
	for _, i2 := range x {
		y = append(y, i2)
	}
	return y
}

func sortString(s string)string{
	s1 := strings.Split(s,"")
	sort.Strings(s1)
	return strings.Join(s1,"")
}