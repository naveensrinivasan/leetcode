func tribonacci(n int) int {
    a,b,c,d := 0,1,1,0
    if n == 1 || n== 2{
        return 1
    }else if n == 0{
        return 0
    }
    for i:=3;i<=n;i++{
        d = a + b + c
        a,b,c = b,c,d
    }
    return d
}
