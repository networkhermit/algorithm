package main

import "fun/vac/util/TestRunner"

func testParseTest() {
    for i := 0; i < 10; i++ {
        if (i & 1) == 0 {
            TestRunner.ParseTest(false)
        } else {
            TestRunner.ParseTest(true)
        }
    }
}

func main() {
    testParseTest()
}
