func xorOperation(n int, start int) int {
    sum := 0
    for i:=0;i<n;i++{
        sum = sum ^ (start+(2*i))
    }
    return sum
}