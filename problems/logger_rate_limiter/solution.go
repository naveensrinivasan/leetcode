type Logger struct {
    M map[string]int
}


/** Initialize your data structure here. */
func Constructor() Logger {
    return Logger{make(map[string]int)}
}


/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
        If this method returns false, the message will not be printed.
        The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    if v,ok := this.M[message];!ok{
        this.M[message]=timestamp
        return true
    }else{
        if timestamp >= v+10{
            this.M[message]=timestamp
            return true
        }
        return false
    }
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */