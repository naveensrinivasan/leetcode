func removeDuplicates(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    prev,counter := nums[0],1
    for i:=1;i<len(nums);i++{
        if prev != nums[i]{
            nums[counter] = nums[i]
            counter++
        }
        prev= nums[i]
    }
    return counter
}