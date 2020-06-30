import (

	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	return sortString(s) == sortString(t)
}

func sortString(s string)string{
	s1 := strings.Split(s,"")
	sort.Strings(s1)
	return strings.Join(s1,"")
}