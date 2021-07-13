func sumEvenAfterQueries(nums []int, queries [][]int) []int {
    result :=[]int{}
    sum := 0
    for _, i := range nums{
        if i%2==0{
            sum+=i
        }
    }
    for _,item := range queries{
        isEven := nums[item[1]] %2 ==0
        before := nums[item[1]]
        nums[item[1]]+= item[0]
        isEvenAfterSum := nums[item[1]]%2==0
        
        if !isEven && isEvenAfterSum{
            sum = sum+nums[item[1]]
            result = append(result,sum)
        }else if isEven && isEvenAfterSum{
            sum += item[0]
            result = append(result,sum)
        }else if !isEven && !isEvenAfterSum{
            result = append(result,sum)
        }else if isEven &&  !isEvenAfterSum{
            sum -= before
            result = append(result,sum)
        }
        
    }
    return result
}