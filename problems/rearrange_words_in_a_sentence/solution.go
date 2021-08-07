func arrangeWords(s string) string {
    arr := strings.Fields(s)
    sort.SliceStable(arr,func(i,j int)bool{
        return len(arr[i]) < len(arr[j])
    })
    result := ""
    for i,item := range arr{
        if i == 0{
            result += strings.Title(strings.ToLower(item))
        }else{
            result += " " + strings.ToLower(item)
        }
    }
    return result
    
}