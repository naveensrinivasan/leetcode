func countTriples(n int) int {
    counter := 0
    m := make(map[int]int)
    for i:=1;i<=n;i++{
        m[i*i] = i
    }
    for i:=1;i<n;i++{
        for j:=i;j<n;j++{
            if _,ok := m[i*i+j*j];ok{
                counter+=2
            }
        }
    }
    return counter
}