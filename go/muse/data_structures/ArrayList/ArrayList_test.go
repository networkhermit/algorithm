package ArrayList

import "testing"

func TestArrayList(t *testing.T) {
	size := 8192

	list := New(0)

	for i := 1; i <= size; i++ {
		list.Append(i)
	}

	list.Shrink()

	if list.Size() != size {
		t.FailNow()
	}

	if list.Capacity() != size {
		t.FailNow()
	}

	for i := 0; i < size; i++ {
		if list.Get(i) != i+1 {
			t.FailNow()
		}
	}

	for i := 0; i < size; i++ {
		list.Set(i, size-i)
	}

	for i := 0; i < size; i++ {
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

	for i := 0; i < size; i++ {
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

	list.Shrink()

	if !list.IsEmpty() {
		t.FailNow()
	}

	if list.Capacity() != 0 {
		t.FailNow()
	}
}
