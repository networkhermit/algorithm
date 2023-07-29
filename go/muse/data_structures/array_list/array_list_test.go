package array_list

import (
	"testing"

	"muse/data_structures/tests"
)

func TestArrayList(t *testing.T) {
	t.Run("ArrayList", tests.ListDerive(New))
}
