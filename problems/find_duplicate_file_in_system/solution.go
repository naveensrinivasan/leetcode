func findDuplicate(paths []string) [][]string {
    m := make(map[string][]string)
    for _,item := range paths{
        data := strings.Fields(item)
        for i:=1;i<len(data);i++{
            x := strings.Split(data[i],"(")
            f := data[0]+"/"+x[0]
            content := strings.Trim(x[1],")")
            if v,ok := m[content];!ok{
                m[content] = []string{f}
            }else{
                v = append(v,f)
                m[content] = v
            }
        }
    }
    result := [][]string{}
    for _, v := range m{
        if len(v) > 1{
            result = append(result,v)
        }
    }
    return result
}