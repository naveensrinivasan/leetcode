func reverseString(s []byte)  {
    left,right := 0, len(s) -1
    for i:=0;i<len(s)/2;i++{
        s[left+i], s[right-i] =  s[right-i] , s[left+i]
    }
}