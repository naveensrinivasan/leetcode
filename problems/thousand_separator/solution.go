func thousandSeparator(n int) string {
    if n < 1000{
        return strconv.Itoa(n)
    }
    x:= strconv.Itoa(n)
    result := ""
    for i,y := range strings.Split(Reverse(x),""){
        if i != 0 && (i)%3 == 0{
          result += "."+y            
        }else{
            result += y
        } 
    }
    return Reverse(result)
}
// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}