func sortArrayByParityII(nums []int) []int {
    result := make([]int,len(nums))
    even,odd:= 0,1
    for i:=0;i<len(nums);i++{
        if nums[i]%2==0{
            result[even] = nums[i]
            even += 2
        }else{
            result[odd] = nums[i]
            odd+=2
        }
    }
    return result
}