func findRestaurant(list1 []string, list2 []string) []string {
    m := make(map[string]int,len(list1))
    result := []string{}
    min := 2001
    for i, word := range list1{
        m[word] = i
    }
    for i, word := range list2{
        if v,ok := m[word];ok{
            if v+i == min{
                result = append(result,word)
            }else if v+i < min{
                min = v+i
                result = []string{word}
            }else{
                fmt.Println(v+i)
            }
        }
    }
    return result
}