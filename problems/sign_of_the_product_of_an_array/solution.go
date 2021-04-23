func arraySign(nums []int) int {
    negatives := 0
    for _, i := range nums{
        if i == 0{
            return 0
        }else if i < 0{
            negatives++
        }
    }
    if negatives %2 == 0{
        return 1
    }
    
    return -1
}