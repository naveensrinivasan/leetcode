func minSteps(s string, t string) int {
    m:= make(map[string]int)
    k:= make(map[string]int)
    for i:=0;i<len(t);i++{
        m[string(s[i])]++
        k[string(t[i])]++
    }
    result :=0
    for key,value := range m{
        if v,ok := k[key];!ok{
            result +=value
        }else{
            if value > v{
                result+= (value -v)
            }
        }
    }
    return result
}