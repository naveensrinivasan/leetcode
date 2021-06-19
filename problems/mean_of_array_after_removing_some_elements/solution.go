func trimMean(arr []int) float64 {
    cut := int(float64(len(arr))*float64(.05))
    sort.Ints(arr)
    arr = arr[cut:len(arr)-cut]
    sum := float64(0)
    for _, i := range arr{
        sum+= float64(i)
    }
    return sum/float64(len(arr))
}