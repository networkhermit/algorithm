package singly_linked_list

import (
	"testing"

	"muse/data_structures/tests"
)

func TestSinglyLinkedList(t *testing.T) {
	t.Run("SinglyLinkedList", tests.ListDerive(New))
}
