func moveZeroes(nums []int)  {
    counter := 0
    for i:=0;i<len(nums);i++{
        if nums[i] != 0{
            nums[counter] = nums[i]
            counter++
        }
    }
    for i:=counter;i<len(nums);i++{
        nums[i] = 0
    }
}