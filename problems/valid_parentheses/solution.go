func isValid(s string) bool {
    stack := []string{}
    for _, i := range s{
        k := string(i)
        if len(stack) > 0 && k == ")" {
            if stack[len(stack)-1] == "(" {
                stack = stack[:len(stack)-1]
            }else{
                return false
            }
        }else if len(stack) > 0 && k == "}" {
            if stack[len(stack)-1] == "{" {
                stack = stack[:len(stack)-1]
            }else{
                return false
            }
        }else if len(stack) > 0 && k == "]"{
            if stack[len(stack)-1] == "[" {
                stack = stack[:len(stack)-1]
            }else{
                return false
            }
        }else{
            stack = append(stack,k)
        }
    
        
    }
   
    return len(stack) == 0
    
}