func isMajorityElement(nums []int, target int) bool {
    counter := 0
    for _,j := range nums{
        if j == target{
           counter++ 
        }
        
    }
    
    return counter > (len(nums)/2)
    
}