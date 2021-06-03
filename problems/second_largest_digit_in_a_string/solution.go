func secondHighest(s string) int {
    l := []int{}
    m := make(map[int]int)
    for i:=0;i<len(s);i++{
        if s[i] >=48 && s[i]<=57{
            x,_ := strconv.Atoi(string(s[i]))        
            if _,ok:= m[x];!ok{
                m[x]=x
                l = append(l,x) 
            }
        }
    }
    
    sort.Slice(l,func(i,j int)bool{
                    return l[i] > l [j]
                })
    
    if len(l) >=2 {
        return l[1]
    }
    
    return -1
}