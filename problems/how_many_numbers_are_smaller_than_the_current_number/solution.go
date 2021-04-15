func smallerNumbersThanCurrent(nums []int) []int {
    result := []int{}
    for i := 0; i< len(nums); i++{
        counter := 0
        for j:= 0; j<len(nums);j++{
            if nums[j] < nums[i]{
                counter++
            } 
        }
        result = append(result,counter)
    }
    return result
}