package main

import (
    "fun/vac/util/RandomFactory"
    "fun/vac/util/TestRunner"
)

func testGenerateEven() bool {
    RandomFactory.Launch()

    var value int
    for i := 0; i < 8192; i++ {
        value = RandomFactory.GenerateEven()
        if (value & 1) != 0 {
            return false
        }
    }

    return true
}

func testGenerateOdd() bool {
    RandomFactory.Launch()

    var value int
    for i := 0; i < 8192; i++ {
        value = RandomFactory.GenerateOdd()
        if (value & 1) == 0 {
            return false
        }
    }

    return true
}

func main() {

    TestRunner.ParseTest(testGenerateEven())

    TestRunner.ParseTest(testGenerateOdd())
}
