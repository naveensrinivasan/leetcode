func countOdds(low int, high int) int {
    oddl:= low%2==1
    oddh:= high%2==1
    
    if (oddl && oddh) || (oddl || oddh){
        return ((high-low)/2)+1
    }
    return ((high-low)/2)
}