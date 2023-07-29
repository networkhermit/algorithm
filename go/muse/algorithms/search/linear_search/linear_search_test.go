package linear_search

import (
	"testing"

	"muse/algorithms/search/tests"
)

func TestLinearSearch(t *testing.T) {
	t.Run("LinearSearch", tests.Derive(Find))
}
