func getModifiedArray(length int, updates [][]int) []int {
    arr := make([]int,length)
    for _,item := range updates{
        for i:=item[0];i<=item[1];i++{
            arr[i]+=item[2]
        } 
    }
    return arr
}