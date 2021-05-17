

func commonChars(A []string) []string {
    m := []map[string]int{}
    result := []string{}
    maxItems := make(map[string]int)
    
    
    for _,word := range A{
        m2 := make(map[string]int)
        for _,l := range word{
            m2[string(l)]++
        }
        if len(m2) > len(maxItems) {
            maxItems = m2
        }
        
        m = append(m,m2)
        
    }
    
    fmt.Println(m)
    
     for k := range maxItems{
         exists := true
         min := 1000000
         for i:= 0;i<len(m);i++{
             if value,ok:=m[i][k];ok{
                 if value < min{
                     min = value
                 }
             }else{
                 exists = false
                 break
             } 
         }
         if exists{
             for i:= 0;i<min;i++{
                result = append(result,k)    
             }
             
         }
     }
    fmt.Println(result)
    return result
}