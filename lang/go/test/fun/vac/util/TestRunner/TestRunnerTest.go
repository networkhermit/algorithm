package main

import "fun/vac/util/TestRunner"

func testParseTest() {
    for i := 0; i < 10; i++ {
        TestRunner.ParseTest((i & 1) != 0)
    }
}

func main() {
    testParseTest()
}
