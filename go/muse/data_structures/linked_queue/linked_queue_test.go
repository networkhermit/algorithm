package linked_queue

import (
	"testing"

	"muse/data_structures/tests"
)

func TestLinkedQueue(t *testing.T) {
	t.Run("LinkedQueue", tests.QueueDerive(New[int]))
}
