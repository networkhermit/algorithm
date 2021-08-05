from muse.util import RandomFactory
from muse.util import TestRunner

def testIntegerN() -> bool:
    RandomFactory.seed()

    value = 0
    for _ in range(8192):
        if RandomFactory.integerN(0, 0) != 0:
            return False

        if RandomFactory.integerN(1, 1) != 1:
            return False

        value = RandomFactory.integerN(0, 1)
        if value < 0 or value > 1:
            return False

        value = RandomFactory.integerN(100, 10000)
        if RandomFactory.integerN(value, value) != value:
            return False
        if value < 100 or value > 10000:
            return False

    return True

def testGenerateEven() -> bool:
    RandomFactory.seed()

    for _ in range(8192):
        if (RandomFactory.generateEven() & 1) != 0:
            return False

    return True

def testGenerateOdd() -> bool:
    RandomFactory.seed()

    for _ in range(8192):
        if (RandomFactory.generateOdd() & 1) == 0:
            return False

    return True

def main() -> None:
    TestRunner.pick(testIntegerN)

    TestRunner.pick(testGenerateEven)

    TestRunner.pick(testGenerateOdd)

if __name__ == "__main__":
    main()
