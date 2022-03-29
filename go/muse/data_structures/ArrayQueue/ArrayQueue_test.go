package ArrayQueue

import "testing"

func TestArrayQueue(t *testing.T) {
	size := 8192

	queue := New(0)

	for i := 1; i <= size; i++ {
		queue.Enqueue(i)
	}

	queue.Shrink()

	if queue.Size() != size {
		t.FailNow()
	}

	if queue.Capacity() != size {
		t.FailNow()
	}

	for i := 1; i <= size; i++ {
		if queue.Peek() != i {
			t.FailNow()
		}
		queue.Dequeue()
	}

	queue.Shrink()

	if !queue.IsEmpty() {
		t.FailNow()
	}

	if queue.Capacity() != 0 {
		t.FailNow()
	}
}
