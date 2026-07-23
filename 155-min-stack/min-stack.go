type Pair struct {
    Val int
    Min int
}

type MinStack struct {
    stack []Pair
}


func Constructor() MinStack {
    return MinStack{
        stack: []Pair{},
    }
}


func (this *MinStack) Push(value int)  {
    if len(this.stack) == 0 {
        this.stack = append(this.stack, Pair{Val: value, Min: value})
        return
    }
    curr_min := this.stack[len(this.stack)-1].Min
    if value < curr_min {
        curr_min = value
    }
    this.stack = append(this.stack, Pair{Val: value, Min: curr_min})
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1].Val
}


func (this *MinStack) GetMin() int {
    return this.stack[len(this.stack)-1].Min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */