package main

import (
    "muse/util/RandomFactory"
    "muse/util/TestRunner"
)

func testIntegerN() bool {
    RandomFactory.Seed()

    var value int
    for i := 0; i < 8192; i++ {
        if RandomFactory.IntegerN(0, 0) != 0 {
            return false
        }

        if RandomFactory.IntegerN(1, 1) != 1 {
            return false
        }

        value = RandomFactory.IntegerN(0, 1)
        if value < 0 || value > 1 {
            return false
        }

        value = RandomFactory.IntegerN(100, 10000)
        if RandomFactory.IntegerN(value, value) != value {
            return false
        }
        if value < 100 || value > 10000 {
            return false
        }
    }

    return true
}

func testGenerateEven() bool {
    RandomFactory.Seed()

    for i := 0; i < 8192; i++ {
        if (RandomFactory.GenerateEven() & 1) != 0 {
            return false
        }
    }

    return true
}

func testGenerateOdd() bool {
    RandomFactory.Seed()

    for i := 0; i < 8192; i++ {
        if (RandomFactory.GenerateOdd() & 1) == 0 {
            return false
        }
    }

    return true
}

func main() {
    TestRunner.ParseTest(testIntegerN())

    TestRunner.ParseTest(testGenerateEven())

    TestRunner.ParseTest(testGenerateOdd())
}
