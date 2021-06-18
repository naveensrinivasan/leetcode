func highFive(items [][]int) [][]int {
    m := make(map[int][]int)
    for _,student := range items{
        if v,ok := m[student[0]];!ok{
            l := []int{student[1]}
            m[student[0]] = l
        }else{
            v = append(v,student[1])
            m[student[0]] = v
        }
    }
    
    result :=[][]int{}
    for k,v := range m{
        sort.Slice(v,func(i,j int)bool{
            return v[i] > v[j]
        })
        sum :=0
        for i:=0;i<5;i++{
            sum+= v[i]
        }
        result = append(result,[]int{k,sum/5})
    }
    sort.Slice(result,func(i,j int)bool{
        return result[i][0] < result[j][0]
    })
    return result
}