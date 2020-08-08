func transpose(A [][]int) [][]int {
	result := [][]int{}
	for col := 0; col < len(A[0]); col++ {
		mrow := []int{}
		for row := 0; row < len(A); row++ {
			mrow = append(mrow, A[row][col])
		}
		result = append(result, mrow)
	}
	return result

}
