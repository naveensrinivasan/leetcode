func countPrimeSetBits(left int, right int) int {
    result := 0
    for i:=left;i<=right;i++{
        count :=0
        x := i    
        for x > 0{
            if x&1==1 {
                count++
            }
            x >>= 1
        }
        if isPrime(count){
            result++
        }
    }
    return result
}

func isPrime(y int)bool{
    if y == 1{
        return false
    }
    if y == 2{
        return true
    }
    x := int(math.Sqrt(float64(y)))
    for i:=2;i<x+1;i++{
        if y%i == 0{
            return false
        }
    }
    return true
}