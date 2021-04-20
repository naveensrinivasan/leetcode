func twoSum(numbers []int, target int) []int {
    
    for index ,_ := range numbers{
        
    
    left,right:= index, len(numbers)-1
    for left<right && right < len(numbers){
        sum := numbers[left] + numbers[right]
        if sum == target{
            return []int{left+1,right+1}
        }else if sum > target{
            right--
        }else{
            left++
        }
    }
    }
    return []int{}
}