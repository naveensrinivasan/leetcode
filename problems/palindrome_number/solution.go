func isPalindrome(x int) bool {
	pop := 0
	reversed := 0
	y := x
	for y != 0 {
		pop = y % 10
		reversed = reversed*10 + pop
		y /= 10
	}
	fmt.Println(reversed, x)
	return x == int(math.Abs(float64(reversed)))
}
