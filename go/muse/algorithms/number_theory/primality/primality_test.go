package primality

import (
	"testing"

	"muse/algorithms/number_theory/primality/tests"
)

func TestPrimality(t *testing.T) {
	t.Parallel()

	t.Run("IsPrime", tests.Derive(IsPrime, tests.Prime))

	t.Run("IsComposite", tests.Derive(IsComposite, tests.Composite))
}
