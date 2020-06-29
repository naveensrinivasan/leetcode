func shuffle(nums []int, n int) []int {
    var x []int
	for i:=0;i<n;i++  {
		x = append(x,nums[i],nums[i+(n)])
	}
	return x
}