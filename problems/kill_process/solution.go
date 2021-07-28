func killProcess(pid []int, ppid []int, kill int) []int {
    m := make(map[int][]int)
    for i:=0;i<len(pid);i++{
        if _,ok := m[pid[i]];!ok{
            m[pid[i]] = []int{}
        }
        if v,ok := m[ppid[i]];!ok{
            m[ppid[i]] = []int{pid[i]}
        }else{
            v = append(v,pid[i])
            m[ppid[i]] =v
        }
    }
    stack :=m[kill]
    result := []int{kill}
    for {
        f := []int{}
        for len(stack) > 0{
            f = append(f,m[stack[0]]...)
            result = append(result,stack[0])
            stack = stack[1:]
        }
        if len(f) > 0{
            stack = append(stack,f...)
        }else{
            break
        }
    }
    return result
}


