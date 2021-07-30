func canAttendMeetings(arr [][]int) bool {
    sort.Slice(arr,func(i,j int)bool{
        return arr[i][0] < arr[j][0]
    })

    for i:=0;i<len(arr)-1;i++ {
        if arr[i+1][0] < arr[i][1]{
            return false
        }
    }
    return true
}