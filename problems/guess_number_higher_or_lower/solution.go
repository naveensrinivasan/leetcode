/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l,r,ans := 1,n,0
    
    for l<=r{
        mid := (l+r)/2
        x := guess(mid)
        if x ==0{
            return mid
        }
        
        if x > 0{
            l = mid +1
            ans = l
        }else{
            r = mid -1
        }
    }
    return ans
}