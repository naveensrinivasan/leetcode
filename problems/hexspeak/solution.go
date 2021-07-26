func toHexspeak(num string) string {
    x , _ := strconv.Atoi(num)
    s := strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(strconv.FormatInt(int64(x),16),"0","O"),"1","I"))
    for _,i := range s{
        if unicode.IsDigit(i){
            return "ERROR"
        }
    }
    return s
}