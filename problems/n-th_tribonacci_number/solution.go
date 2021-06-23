func tribonacci(n int) int {
    m := make(map[int]int)
    m[0] = 0
    m[1] = 1
    m[2] = 1
    for i:=3;i<=n;i++{
        m[i] = m[i-2] + m[i-1] + m[i-3]       
    }
    return m[n]
}
