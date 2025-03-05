package doubly_linked_list

import (
	"testing"

	"muse/data_structures/tests"
)

func TestDoublyLinkedList(t *testing.T) {
	t.Run("DoublyLinkedList", tests.ListDerive(New[int]))
}
