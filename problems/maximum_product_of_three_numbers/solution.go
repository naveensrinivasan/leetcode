func maximumProduct(nums []int) int {
    negativeProduct,positiveProduct:= 0,0
    
    if len(nums)==3{
        return nums[0] * nums[1] * nums[2]
    }
    
   sort.Ints(nums)
    
   negativeProduct =nums[0] * nums[1] * nums[len(nums)-1]
    
   sort.Slice(nums,func(i,j int)bool{
        return nums[i] > nums[j]
    })
    
    positiveProduct = nums[0] * nums[1] * nums[2]
    
    if negativeProduct > positiveProduct{
        return negativeProduct
    }
    
    return positiveProduct
}