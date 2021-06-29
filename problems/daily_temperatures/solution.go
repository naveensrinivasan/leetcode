func dailyTemperatures(temperatures []int) []int {
    type F struct{
        K,I int
    }
    result :=make([]int,len(temperatures))
    s:=[]F{}
    for i:=len(temperatures)-1;i>=0;i--{
        v := 0
        counter := 0
        for j:=len(s)-1;j >= 0 && temperatures[i] >= s[j].K;j,counter = j-1,counter+1{}
        s = s[:len(s)-counter]
        if len(s) > 0{
            v = (s[len(s)-1].I) - i
        }
        result[i] =v
        s = append(s,F{K:temperatures[i],I:i})
    }
    return result
}