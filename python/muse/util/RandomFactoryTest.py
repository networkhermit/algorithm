from muse.util import RandomFactory
from muse.util import TestRunner


def testGenIntN() -> bool:
    value = 0
    for _ in range(8192):
        if RandomFactory.genIntN(0, 0) != 0:
            return False

        if RandomFactory.genIntN(1, 1) != 1:
            return False

        value = RandomFactory.genIntN(0, 1)
        if value < 0 or value > 1:
            return False

        value = RandomFactory.genIntN(100, 10000)
        if RandomFactory.genIntN(value, value) != value:
            return False
        if value < 100 or value > 10000:
            return False

    return True


def testGenEven() -> bool:
    for _ in range(8192):
        if (RandomFactory.genEven() & 1) != 0:
            return False

    return True


def testGenOdd() -> bool:
    for _ in range(8192):
        if (RandomFactory.genOdd() & 1) == 0:
            return False

    return True


def main() -> None:
    TestRunner.pick(testGenIntN)

    TestRunner.pick(testGenEven)

    TestRunner.pick(testGenOdd)


if __name__ == "__main__":
    main()
