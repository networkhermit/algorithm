package tests

import (
	"testing"

	"muse/data_structures"
)

func ListDerive[L data_structures.List](factory func() L) func(t *testing.T) {
	return func(t *testing.T) {
		size := 8192

		list := factory()

		for i := 1; i <= size; i++ {
			list.Append(i)
		}

		if o, ok := any(list).(data_structures.Resizable); ok {
			o.Shrink()
		}

		if list.Size() != size {
			t.FailNow()
		}

		if o, ok := any(list).(data_structures.Resizable); ok {
			if o.Capacity() != size {
				t.FailNow()
			}
		}

		for i := range size {
			if list.Get(i) != i+1 {
				t.FailNow()
			}
		}

		for i := range size {
			list.Set(i, size-i)
		}

		for i := range size {
			if list.Get(i) != size-i {
				t.FailNow()
			}
		}

		for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
			x := list.Get(i)
			y := list.Get(j)

			list.Remove(i)
			list.Insert(i, y)
			list.Remove(j)
			list.Insert(j, x)
		}

		for i := range size {
			if list.Get(i) != i+1 {
				t.FailNow()
			}
		}

		for i := size; i >= 1; i-- {
			if list.Back() != i {
				t.FailNow()
			}
			list.Eject()
		}

		if o, ok := any(list).(data_structures.Resizable); ok {
			o.Shrink()
		}

		if !list.IsEmpty() {
			t.FailNow()
		}

		if o, ok := any(list).(data_structures.Resizable); ok {
			if o.Capacity() != 0 {
				t.FailNow()
			}
		}
	}
}

func QueueDerive[Q data_structures.Queue](factory func() Q) func(t *testing.T) {
	return func(t *testing.T) {
		size := 8192

		queue := factory()

		for i := 1; i <= size; i++ {
			queue.Enqueue(i)
		}

		if o, ok := any(queue).(data_structures.Resizable); ok {
			o.Shrink()
		}

		if queue.Size() != size {
			t.FailNow()
		}

		if o, ok := any(queue).(data_structures.Resizable); ok {
			if o.Capacity() != size {
				t.FailNow()
			}
		}

		for i := 1; i <= size; i++ {
			if queue.Peek() != i {
				t.FailNow()
			}
			queue.Dequeue()
		}

		if o, ok := any(queue).(data_structures.Resizable); ok {
			o.Shrink()
		}

		if !queue.IsEmpty() {
			t.FailNow()
		}

		if o, ok := any(queue).(data_structures.Resizable); ok {
			if o.Capacity() != 0 {
				t.FailNow()
			}
		}
	}
}

func StackDerive[S data_structures.Stack](factory func() S) func(t *testing.T) {
	return func(t *testing.T) {
		size := 8192

		stack := factory()

		for i := 1; i <= size; i++ {
			stack.Push(i)
		}

		if o, ok := any(stack).(data_structures.Resizable); ok {
			o.Shrink()
		}

		if stack.Size() != size {
			t.FailNow()
		}

		if o, ok := any(stack).(data_structures.Resizable); ok {
			if o.Capacity() != size {
				t.FailNow()
			}
		}

		for i := size; i > 0; i-- {
			if stack.Peek() != i {
				t.FailNow()
			}
			stack.Pop()
		}

		if o, ok := any(stack).(data_structures.Resizable); ok {
			o.Shrink()
		}

		if !stack.IsEmpty() {
			t.FailNow()
		}

		if o, ok := any(stack).(data_structures.Resizable); ok {
			if o.Capacity() != 0 {
				t.FailNow()
			}
		}
	}
}
