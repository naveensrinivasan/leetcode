func summaryRanges(nums []int) []string {
    if len(nums) == 0{
        return []string{}
    }
    min,max := nums[0],nums[0]
    result := []string{}
    
    if len(nums) ==1{
        return  []string{strconv.Itoa(nums[0])}
    }
    for i:= 1;i<len(nums);i++{
        if nums[i] == (nums[i-1]+1){
            max = nums[i]
        }else{
            if max > min {
                result = append(result,strconv.Itoa(min) + "->" + strconv.Itoa(max)) 
            }else{
                result = append(result,strconv.Itoa(min))    
            }
            min,max = nums[i] , nums[i]  
        }
        
    }
    if max > min {
                result = append(result,strconv.Itoa(min) + "->" + strconv.Itoa(max)) 
            }else{
                result = append(result,strconv.Itoa(min))    
            }
    return result
}