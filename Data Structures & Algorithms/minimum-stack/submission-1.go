type MinStack struct {
    stack    []int
    minStack []int
}

// Constructor initializes the MinStack data structure.
func Constructor() MinStack {
    return MinStack{
        // Pre-allocate a small capacity to reduce initial array reallocations.
        stack:    make([]int, 0, 16),
        minStack: make([]int, 0, 16),
    }
}

// Push adds the element val onto the stack.
func (this *MinStack) Push(val int) {
    this.stack = append(this.stack, val)

	// If minStack is empty or val is smaller/equal to current min, push to minStack
    if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
        this.minStack = append(this.minStack, val)
    }
}

// Pop removes the element on the top of the stack.
func (this *MinStack) Pop() {
    if len(this.stack) == 0 {
        return
    }

    // If the element being popped is the current minimum, pop it from minStack too.
    topVal := this.stack[len(this.stack)-1]
    if topVal == this.minStack[len(this.minStack)-1] {
        this.minStack = this.minStack[:len(this.minStack)-1]
    }

    // Pop from the main stack.
    this.stack = this.stack[:len(this.stack)-1]
}

// Top gets the top element of the stack.
func (this *MinStack) Top() int {
    if len(this.stack) == 0 {
        panic("Top() called on an empty stack") 
    }
    return this.stack[len(this.stack)-1]
}

// GetMin retrieves the minimum element in the stack.
func (this *MinStack) GetMin() int {
    if len(this.minStack) == 0 {
        panic("GetMin() called on an empty stack")
    }
    return this.minStack[len(this.minStack)-1]
}