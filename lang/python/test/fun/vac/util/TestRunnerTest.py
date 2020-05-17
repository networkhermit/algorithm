from fun.vac.util import TestRunner

def testParseTest() -> None:
    for i in range(10):
        if (i & 1) == 0:
            TestRunner.parseTest(False)
        else:
            TestRunner.parseTest(True)

def main() -> None:
    testParseTest()

if __name__ == "__main__":
    main()
