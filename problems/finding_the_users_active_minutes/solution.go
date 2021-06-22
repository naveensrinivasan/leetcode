func findingUsersActiveMinutes(logs [][]int, k int) []int {
    m := make(map[int]map[int]int)
    for _, item := range logs {
        if v,ok := m[item[0]];!ok{
            x := make(map[int]int)
            x[item[1]] = item[1]
            m[item[0]] = x
        }else{
            if _, ok := v[item[1]];!ok{
                v[item[1]] = item[1]
            }
            m[item[0]] = v
        }
    }
    x := make(map[int]int)
    for _,v := range m{
        x[len(v)]++
    }
    
    result := []int{}
    for i:=1;i<=k;i++{
        if v,ok:= x[i];ok{
            result = append(result,v)
        }else{
            result = append(result,0)
        }
    }
    return result
}