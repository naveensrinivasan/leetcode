func interpret(command string) string {
    result := ""
    i:= 0
    for i < len(command){
        if command[i] == 'G'{
            result += "G"
            i+=1
        }else if command[i] == '(' && command[i+1] == ')'{
            result += "o"
            i +=2
            
    }else {
           result += "al"
           i+=4
        }
    }
    return result
}