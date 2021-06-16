func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    result := 0
    for _, i := range arr1{
        flag := false
        for _, j:= range arr2 {
            if abs(i,j) <= d{
                flag = true
            }
        }
        if !flag{
            result++
        }
    }
    return result
}

func abs(i,j int)int{
    x := i-j
    if x < 0 {
        return x * -1
    }
    return x
}