func isPrefixString(s string, words []string) bool {
    counter := 0
    
    for i,item := range words{
        if i == 0 && len(item) > len(s){
            return false
        }
        if len(s) >= counter+len(item) {
            if s[counter:counter+len(item)] != item{
                return false
            }
        }else{
            break
        }
        counter += len(item)
    }
    if counter>= len(s){
        return true
    }
    return false
}