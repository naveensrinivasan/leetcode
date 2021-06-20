func minOperations(logs []string) int {
    l := []string{}
    for _, path := range logs{
        switch path{
            case "../":
            if len(l) > 0{
                l = l[:len(l)-1]
            }
            case "./":
            default:
            l = append(l,path)
        }
    }
    return len(l)
}