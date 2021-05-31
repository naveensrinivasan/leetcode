func findOcurrences(text string, first string, second string) []string {
    result := []string{}
    l := strings.Split(text," ")
    for i:= 1;i<len(l);i++{
        if l[i-1] == first && l[i]== second && i+1 < len(l){
            result = append(result,l[i+1])
        }
    }
    return result
}