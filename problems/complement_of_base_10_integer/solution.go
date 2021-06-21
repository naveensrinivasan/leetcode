func bitwiseComplement(N int) int {
    counter := 0
    f := N
    x:= 0
    for N > 0{
        counter++
        N >>= 1
    }
    for i:=0;i<counter-1;i++{
        x = x^1
        x <<=1
    }
    x = x^1
    return f ^ x
}