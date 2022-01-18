// Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement Queue using Stacks.
// Memory Usage: 2.3 MB, less than 10.53% of Go online submissions for Implement Queue using Stacks.

type MyQueue struct {
    pushOnlyStack []int
    popOnlyStack  []int
}


func Constructor() MyQueue {
    myQueue := new(MyQueue)
    myQueue.pushOnlyStack = make([]int, 0)
    myQueue.popOnlyStack = make([]int, 0)
    return *myQueue
}


func (this *MyQueue) Push(x int)  {
    if !this.isPopOnlyEmpty() {
        this.transferElemsToPushOnly()
    }
    this.pushOnlyStack = append(this.pushOnlyStack, x)
}


func (this *MyQueue) Pop() int {
    if !this.isPushOnlyEmpty() {
        this.transferElemsToPopOnly()
    }
    ele := this.popOnlyStack[len(this.popOnlyStack)-1]
    this.popOnlyStack = this.popOnlyStack[:len(this.popOnlyStack)-1]
    return ele
}


func (this *MyQueue) Peek() int {
    if !this.isPushOnlyEmpty() {
        this.transferElemsToPopOnly()
    }
    ele := this.popOnlyStack[len(this.popOnlyStack)-1]
    return ele
}


func (this *MyQueue) Empty() bool {
    return len(this.pushOnlyStack) == 0 && len(this.popOnlyStack) == 0
}

func (this *MyQueue) isPushOnlyEmpty() bool {
    return len(this.pushOnlyStack) == 0
}

func (this *MyQueue) isPopOnlyEmpty() bool {
    return len(this.popOnlyStack) == 0
}

func (this *MyQueue) transferElemsToPushOnly() {
    for _, _ = range this.popOnlyStack {
        this.pushOnlyStack = append(this.pushOnlyStack, this.popOnlyStack[len(this.popOnlyStack)-1])
        this.popOnlyStack = this.popOnlyStack[:len(this.popOnlyStack)-1]
    }
}

func (this *MyQueue) transferElemsToPopOnly() {
    for _, _ = range this.pushOnlyStack {
        this.popOnlyStack = append(this.popOnlyStack, this.pushOnlyStack[len(this.pushOnlyStack)-1])
        this.pushOnlyStack = this.pushOnlyStack[:len(this.pushOnlyStack)-1]
    }
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */