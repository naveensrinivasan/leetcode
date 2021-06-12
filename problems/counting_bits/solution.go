func countBits(n int) []int {
    result := []int{}
    for i:=0;i<=n;i++{
        counter:=0
        x :=i
        for x>0{
            counter += x & 1 
            x = x >> 1
        }
        result = append(result,counter)
    }
    return result
}