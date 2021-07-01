func decompressRLElist(nums []int) []int {
    result := []int{}
    for i:=0;i<len(nums);i= i+2{
        for j:=0;j<nums[i];j++{
            result = append(result,nums[i+1])
        }
    }  
    return result
}