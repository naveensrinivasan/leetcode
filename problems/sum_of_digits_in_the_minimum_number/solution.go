func sumOfDigits(nums []int) int {
    sort.Ints(nums)
    i := nums[0]
    if i < 10{
        if i %2 == 1{
            return 0
        }
        return 1
    }else{
        sum := 0
        for i > 0{
            sum += i %10
            i /=10
        }
        if sum %2 == 1{
            return 0
        }
        return 1
    } 
    
    return 90
}