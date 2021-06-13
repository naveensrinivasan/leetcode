type BrowserHistory struct {
    L  []string
    pos  int
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{[]string{homepage},0}
}


func (this *BrowserHistory) Visit(url string)  {
    if this.pos < len(this.L)-1{
        this.L = this.L[:this.pos+1]
    }
    this.L = append(this.L,url)
    this.pos++
}


func (this *BrowserHistory) Back(steps int) string {
    if this.pos - steps < 0{
        this.pos = 0
        return this.L[0]
    }
    this.pos -= steps
    return this.L[this.pos]
}


func (this *BrowserHistory) Forward(steps int) string {
    if this.pos + steps > len(this.L)-1{
        this.pos = len(this.L) -1
        return this.L[len(this.L)-1]
    }
    this.pos += steps 
    return this.L[this.pos]
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */