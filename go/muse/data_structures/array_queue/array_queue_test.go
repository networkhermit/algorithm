package array_queue

import (
	"testing"

	"muse/data_structures/tests"
)

func TestArrayQueue(t *testing.T) {
	t.Run("ArrayQueue", tests.QueueDerive(New))
}
