package ArrayStack

import "testing"

func TestArrayStack(t *testing.T) {
	size := 8192

	stack := New(0)

	for i := 1; i <= size; i++ {
		stack.Push(i)
	}

	stack.Shrink()

	if stack.Size() != size {
		t.FailNow()
	}

	if stack.Capacity() != size {
		t.FailNow()
	}

	for i := size; i > 0; i-- {
		if stack.Peek() != i {
			t.FailNow()
		}
		stack.Pop()
	}

	stack.Shrink()

	if !stack.IsEmpty() {
		t.FailNow()
	}

	if stack.Capacity() != 0 {
		t.FailNow()
	}
}
