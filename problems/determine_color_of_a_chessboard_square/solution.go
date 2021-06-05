func squareIsWhite(coordinates string) bool {
    odd := make(map[string]string)
    odd["a"] = "a"
    odd["c"] = "c"
    odd["e"] = "e"
    odd["g"] = "g"
    
    data := strings.Split(coordinates,"")
    a := data[0]
    x,_ := strconv.Atoi(data[1]) 
    
    if _,ok:= odd[a];ok {
        if x %2 ==1 {
            return false
        } 
        return true
    }
    
    if x%2==1{
        return true
    }
    return false
}