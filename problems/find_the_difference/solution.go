func findTheDifference(s string, t string) byte {
    s= SortString(s)
    t= SortString(t)
    
    for i:= 0;i<len(s);i++{
        if s[i] != t[i]{
            return t[i]
        }
    }
    return t[len(t)-1]
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}