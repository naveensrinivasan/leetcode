func uniqueMorseRepresentations(words []string) int {
    x := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.." }
    m := make(map[string]int)
    for _,word := range words{
        w := ""
        for _, i := range word{
            w+= x[i-'a']
        }
        m[w]++
    }
    return len(m)
}
