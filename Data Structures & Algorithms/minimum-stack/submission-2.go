type node struct {
    val int
    min int
}

type MinStack struct {
    stack []node
    head  int // Manual index pointer to the top of the stack
}

func Constructor() MinStack {
    // Assuming a constraint of 30,000 operations.
    // Adjust this constant based on the exact problem bounds.
    return MinStack{
        stack: make([]node, 30000),
        head:  -1, // -1 indicates an empty stack
    }
}

func (this *MinStack) Push(val int) {
    this.head++
    min := val
    
    if this.head > 0 {
        prevMin := this.stack[this.head-1].min
        if prevMin < min {
            min = prevMin
        }
    }
    
    // Direct memory write; no append() capacity checks
    this.stack[this.head] = node{val: val, min: min}
}

func (this *MinStack) Pop() {
    if this.head >= 0 {
        this.head-- // Just move the pointer backward; $O(1)$ and 0 allocations
    }
}

func (this *MinStack) Top() int {
    return this.stack[this.head].val
}

func (this *MinStack) GetMin() int {
    return this.stack[this.head].min
}