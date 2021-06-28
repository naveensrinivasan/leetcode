func dailyTemperatures(temperatures []int) []int {
    type F struct{
        K,I int
    }
    result :=make([]int,len(temperatures))
    s:=[]F{}
    for i:=len(temperatures)-1;i>=0;i--{
        v := 0
        for len(s) > 0 && temperatures[i] >= s[len(s)-1].K{
            s = s[:len(s)-1]
        }
        if len(s) > 0{
            v = (s[len(s)-1].I) - i
        }
        result[i] =v
        s = append(s,F{K:temperatures[i],I:i})
    }
    return result
}