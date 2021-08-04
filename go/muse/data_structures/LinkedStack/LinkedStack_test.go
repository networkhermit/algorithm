package LinkedStack

import "testing"

func TestLinkedStack(t *testing.T) {
    size := 8192

    stack := New()

    for i := 1; i <= size; i++ {
        stack.Push(i)
    }

    if stack.Size() != size {
        t.FailNow()
    }

    for i := size; i > 0; i-- {
        if stack.Peek() != i {
            t.FailNow()
        }
        stack.Pop()
    }

    if !stack.IsEmpty() {
        t.FailNow()
    }
}
