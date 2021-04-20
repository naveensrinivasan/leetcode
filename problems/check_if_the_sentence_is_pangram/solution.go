func checkIfPangram(sentence string) bool {
    m := make([]int,26)
    total := 0
    for _,i := range sentence{
        wasSet := false
        if  m[int(i) - int('a')] == 0{
         m[int(i) - int('a')] =1   
         wasSet = true
        }
         
        if m[int(i) - int('a')] == 1 && wasSet{
            total++
        }
    }  
    return total ==26 
}