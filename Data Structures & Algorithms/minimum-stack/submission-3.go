type MinStack struct {
    stack []int
    min   int
}

func Constructor() MinStack {
    return MinStack{
        stack: make([]int, 0, 16),
    }
}

func (this *MinStack) Push(val int) {
    if len(this.stack) == 0 {
        this.min = val
        this.stack = append(this.stack, val)
        return
    }

    if val >= this.min {
        this.stack = append(this.stack, val)
    } else {
        // The new value is smaller. Encode it to flag the minimum change.
        // NOTE: In Go, `int` is 64-bit on 64-bit architectures, so this equation
        // safely handles standard 32-bit integer limits without underflowing.
        encoded := 2*val - this.min 
        this.stack = append(this.stack, encoded)
        this.min = val // Update the current minimum
    }
}

func (this *MinStack) Pop() {
    topVal := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]

    if topVal < this.min {
        // The popped element is an encoded flag. 
        // Restore the previous minimum using the inverse algebraic equation.
        this.min = 2*this.min - topVal 
    }
}

func (this *MinStack) Top() int {
    topVal := this.stack[len(this.stack)-1]
    if topVal < this.min {
        // If it's encoded, the actual inserted value was the current min.
        return this.min 
    }
    return topVal
}

func (this *MinStack) GetMin() int {
    return this.min
}