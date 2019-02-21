from fun.vac.util import RandomFactory
from fun.vac.util import TestRunner

def testGenerateEven() -> bool:
    RandomFactory.launch()

    value = 0
    for i in range(8192):
        value = RandomFactory.generateEven()
        if (value & 1) != 0:
            return False

    return True

def testGenerateOdd() -> bool:
    RandomFactory.launch()

    value = 0
    for i in range(8192):
        value = RandomFactory.generateOdd()
        if (value & 1) == 0:
            return False

    return True

if __name__ == "__main__":

    TestRunner.parseTest(testGenerateEven())

    TestRunner.parseTest(testGenerateOdd())
