func countPrimes(n int) int {
	count := 0
	for i := 2; i < n; i++ {
		isdivisble := false

		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isdivisble = true
				break
			}
		}
		if !isdivisble {
			//fmt.Println(i)
			count++
		}
	}
	return count

}
