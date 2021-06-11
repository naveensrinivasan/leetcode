func numUniqueEmails(emails []string) int {
    m := make(map[string]string)
    for _, e := range emails{
        data := strings.Split(e,"@")
        person := strings.ReplaceAll(data[0],".","")
        person = strings.Split(person,"+")[0]
        email := person+"@"+data[1]
        if _,ok:= m[email];!ok{
            m[email] = email
        }
    }
    fmt.Println(m)
    return len(m)
}