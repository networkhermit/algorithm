from muse.util import TestRunner

def testParseTest() -> None:
    for i in range(10):
        TestRunner.parseTest((i & 1) != 0)

def main() -> None:
    testParseTest()

if __name__ == "__main__":
    main()
