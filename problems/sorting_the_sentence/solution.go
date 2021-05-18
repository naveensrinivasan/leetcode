func sortSentence(s string) string {
    x:= strings.Split(s," ")
    r := make([]string,len(x))
    
    for _, w := range x{
        i,_ :=strconv.Atoi(string(w[len(w)-1]))
        r[i-1] = w[:len(w)-1]
    }
    fmt.Println(x,r)
    return strings.Join(r, " ")
}