func plusOne(digits []int) []int {
    result := []int{}
    carry := 0
    doIncrement := true
    isFirstTime := true
    for i:=len(digits)-1;i>=0;i--{
        x:= 0
        if doIncrement{
            if isFirstTime {
                x = digits[i] +1 + carry    
                isFirstTime = false
            }else{
                x = digits[i] + carry
            }
            
            doIncrement = false
            if x > 9{
                carry = 1
                doIncrement = true
                result = append([]int{x%10}, result...)    
            }else{
                carry = 0
                result = append([]int{x}, result...)    
            }
            
        }else{
            result = append([]int{digits[i]}, result...)
        }
        
        
    }
    if carry == 1{
        result = append([]int{1}, result...)    
    }
    return result
}