package RandomFactory

import "testing"

func TestIntegerN(t *testing.T) {
    Seed()

    var value int
    for i := 0; i < 8192; i++ {
        if IntegerN(0, 0) != 0 {
            t.FailNow()
        }

        if IntegerN(1, 1) != 1 {
            t.FailNow()
        }

        value = IntegerN(0, 1)
        if value < 0 || value > 1 {
            t.FailNow()
        }

        value = IntegerN(100, 10000)
        if IntegerN(value, value) != value {
            t.FailNow()
        }
        if value < 100 || value > 10000 {
            t.FailNow()
        }
    }
}

func TestGenerateEven(t *testing.T) {
    Seed()

    for i := 0; i < 8192; i++ {
        if (GenerateEven() & 1) != 0 {
            t.FailNow()
        }
    }
}

func TestGenerateOdd(t *testing.T) {
    Seed()

    for i := 0; i < 8192; i++ {
        if (GenerateOdd() & 1) == 0 {
            t.FailNow()
        }
    }
}
