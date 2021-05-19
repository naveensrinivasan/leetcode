func destCity(paths [][]string) string {
    s := make(map[string]string)
    d := make(map[string]string)
    
    for _, path := range paths{
        s[path[0]] = path[0]
        d[path[1]] = path[1]
    }
    for _, source  := range s{
        delete(d,source)
    }
    fmt.Println(s,d)
    for k := range d{
        return k
    }
    
    return ""
}