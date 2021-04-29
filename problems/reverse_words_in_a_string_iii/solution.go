func reverseWords(s string) string {
    result := ""
    for _, item := range strings.Split(s," "){
        
        left,right,mid := 0, len(item)-1, len(item)/2
        
        interim := make([]rune,len(item))
        
        for left<=mid {
            interim[left],interim[right] = rune(item[right]),rune(item[left])
            left++
            right--
        }
        
        result += string(interim) + " "
    }
    return result[:len(result)-1] 
}
