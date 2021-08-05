from muse.util import TestRunner

def testPick() -> None:
    for i in range(10):
        TestRunner.pick(lambda: (i & 1) != 0)

def main() -> None:
    testPick()

if __name__ == "__main__":
    main()
