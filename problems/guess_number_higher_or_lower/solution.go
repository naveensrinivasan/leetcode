/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
// func guess(num int) int {
// 	return 1
// }

//1,2,3,4,5,6,7,8,9,10
// 1+10 /2 = 5

func guessNumber(n int) int {
	low, high := 0, n
	for low <= high {
		mid := (low + high) / 2
		switch guess(mid) {
		case 0:
			return mid
		case -1:
			high = mid - 1
			break
		case 1:
			low = mid + 1
		}
	}
	return n
}
