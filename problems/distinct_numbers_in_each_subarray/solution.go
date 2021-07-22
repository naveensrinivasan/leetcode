func distinctNumbers(nums []int, k int) []int {
    m := make(map[int]int)
    x := 0
    result := []int{}
    for i:=0;i<k;i++{
        if _, ok:= m[nums[i]];!ok{
            m[nums[i]] = 1
            x +=1
        }else{
           m[nums[i]]++
        }
    }
    result = append(result,x)
    for p:=k;p<len(nums);p++{
        if  m[nums[p-k]] == 1{
            delete(m,nums[p-k])            
        }else{
            m[nums[p-k]] --
        }
        m[nums[p]]++
        result = append(result,len(m))
    }
    return result
}