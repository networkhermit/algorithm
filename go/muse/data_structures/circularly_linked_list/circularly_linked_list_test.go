package circularly_linked_list

import (
	"testing"

	"muse/data_structures/tests"
)

func TestCircularlyLinkedList(t *testing.T) {
	t.Run("CircularlyLinkedList", tests.ListDerive(New[int]))
}
