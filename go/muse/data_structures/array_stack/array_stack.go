package array_stack

import "errors"

const DEFAULT_CAPACITY int = 64

type ArrayStack struct {
	data         []int
	logicalSize  int
	physicalSize int
}

func New(physicalSize int) *ArrayStack {
	stack := &ArrayStack{data: nil, logicalSize: 0, physicalSize: DEFAULT_CAPACITY}

	if physicalSize > 1 {
		stack.physicalSize = physicalSize
	}
	stack.data = make([]int, stack.physicalSize)

	return stack
}

func (stack *ArrayStack) Size() int {
	return stack.logicalSize
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.logicalSize == 0
}

func (stack *ArrayStack) Peek() int {
	if stack.logicalSize == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	return stack.data[stack.logicalSize-1]
}

func (stack *ArrayStack) Push(element int) {
	if stack.logicalSize == stack.physicalSize {
		newCapacity := DEFAULT_CAPACITY

		if stack.physicalSize >= DEFAULT_CAPACITY {
			newCapacity = stack.physicalSize + (stack.physicalSize >> 1)
		}

		temp := make([]int, newCapacity)

		for i, length := 0, stack.logicalSize; i < length; i++ {
			temp[i] = stack.data[i]
		}

		stack.data = temp
		stack.physicalSize = newCapacity
	}

	stack.data[stack.logicalSize] = element

	stack.logicalSize++
}

func (stack *ArrayStack) Pop() {
	if stack.logicalSize == 0 {
		panic(errors.New("[PANIC - NoSuchElement]"))
	}

	stack.logicalSize--

	stack.data[stack.logicalSize] = int(0)
}

func (stack *ArrayStack) Capacity() int {
	return stack.physicalSize
}

func (stack *ArrayStack) Shrink() {
	temp := make([]int, stack.logicalSize)

	for i, length := 0, stack.logicalSize; i < length; i++ {
		temp[i] = stack.data[i]
	}

	stack.data = temp
	stack.physicalSize = stack.logicalSize
}
