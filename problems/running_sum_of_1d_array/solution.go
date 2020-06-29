func runningSum(nums []int) []int {
	var x []int
	total := 0
	for _, i2 := range nums {
		total+=i2
		x = append(x,total)
	}
	return x
}