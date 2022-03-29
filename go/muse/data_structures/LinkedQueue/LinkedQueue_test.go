package LinkedQueue

import "testing"

func TestLinkedQueue(t *testing.T) {
	size := 8192

	queue := New()

	for i := 1; i <= size; i++ {
		queue.Enqueue(i)
	}

	if queue.Size() != size {
		t.FailNow()
	}

	for i := 1; i <= size; i++ {
		if queue.Peek() != i {
			t.FailNow()
		}
		queue.Dequeue()
	}

	if !queue.IsEmpty() {
		t.FailNow()
	}
}
