func findComplement(num int) int {
    counter := 0
    f := num
    x:= 0
    for num > 0{
        counter++
        num >>= 1
    }
    for i:=0;i<counter-1;i++{
        x = x^1
        x <<=1
    }
    x = x^1
    return f ^ x
}