func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}
	x := make([][]int, r)
	cols := []int{}
	count := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			cols = append(cols, mat[i][j])
			if len(cols) == c {
				x[count] = cols
				count++
				cols = []int{}
			}
		}
	}
	return x
}