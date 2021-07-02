func findClosestElements(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if abs(arr[i]-x) < abs(arr[j]-x) ||
             (abs(arr[i]-x) == abs(arr[j]-x) && arr[i] < arr[j]) {
			return true
		}
		return false
	})
	y := arr[:k]
	sort.Ints(y)
	return y
}
func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
