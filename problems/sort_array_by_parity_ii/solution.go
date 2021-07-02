func sortArrayByParityII(nums []int) []int {
    i,j,l:=0,1,len(nums)
    for i<l && j<l{
        for i<l && nums[i]%2==0{
           i+=2 
        }
        for j<l && nums[j]%2==1{
           j+=2 
        }
        if (i<l && j <l){
         nums[i], nums[j] = nums[j],nums[i]   
        }
    }
    return nums
}