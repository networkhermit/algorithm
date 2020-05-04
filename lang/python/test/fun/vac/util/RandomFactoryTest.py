from fun.vac.util import RandomFactory
from fun.vac.util import TestRunner

def testIntegerN() -> bool:
    RandomFactory.seed()

    value = 0
    for i in range(8192):
        if RandomFactory.integerN(0, 0) != 0:
            return False

        if RandomFactory.integerN(1, 1) != 1:
            return False

        value = RandomFactory.integerN(0, 1)
        if value < 0 or 1 < value:
            return False

        value = RandomFactory.integerN(100, 10000)
        if RandomFactory.integerN(value, value) != value:
            return False
        if value < 100 or 10000 < value:
            return False

    return True

def testGenerateEven() -> bool:
    RandomFactory.seed()

    for i in range(8192):
        if (RandomFactory.generateEven() & 1) != 0:
            return False

    return True

def testGenerateOdd() -> bool:
    RandomFactory.seed()

    for i in range(8192):
        if (RandomFactory.generateOdd() & 1) == 0:
            return False

    return True

if __name__ == "__main__":

    TestRunner.parseTest(testIntegerN())

    TestRunner.parseTest(testGenerateEven())

    TestRunner.parseTest(testGenerateOdd())
