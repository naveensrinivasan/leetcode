func corpFlightBookings(bookings [][]int, n int) []int {
    result := make([]int,n)
    for _, item := range bookings{
        for i:=item[0];i<=item[1];i++{
            result[i-1]+= item[2]
        }
    }
    return result
}