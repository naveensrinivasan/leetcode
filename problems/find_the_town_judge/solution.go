func findJudge(n int, trust [][]int) int {
    m := make(map[int][]int)
    m1 := make(map[int]bool)
    for _, item := range trust{
        x := item[1]
        if v,ok := m[x];!ok{
            y := []int{}
            y = append(y,item[0])
            m[x] = y
        }else{
            v = append(v,item[0])
            m[x] = v
        }
        m1[item[0]] = true
    }
    x:= 0
    for i:=1;i<=n;i++{
        if _,ok:= m1[i];!ok{
            x = i
            break
        }
    } 
    if x == 0{
        return -1
    }
    fmt.Println(x,m[x])
    if len(m[x]) == n-1{
        return x
    }
   
    return -1
}