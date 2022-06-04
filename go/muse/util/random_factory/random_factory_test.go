package random_factory

import "testing"

func TestGenIntN(t *testing.T) {
	var value int
	for i := 0; i < 8192; i++ {
		if GenIntN(0, 0) != 0 {
			t.FailNow()
		}

		if GenIntN(1, 1) != 1 {
			t.FailNow()
		}

		value = GenIntN(0, 1)
		if value < 0 || value > 1 {
			t.FailNow()
		}

		value = GenIntN(100, 10000)
		if GenIntN(value, value) != value {
			t.FailNow()
		}
		if value < 100 || value > 10000 {
			t.FailNow()
		}
	}
}

func TestGenEven(t *testing.T) {
	for i := 0; i < 8192; i++ {
		if (GenEven() & 1) != 0 {
			t.FailNow()
		}
	}
}

func TestGenOdd(t *testing.T) {
	for i := 0; i < 8192; i++ {
		if (GenOdd() & 1) == 0 {
			t.FailNow()
		}
	}
}
