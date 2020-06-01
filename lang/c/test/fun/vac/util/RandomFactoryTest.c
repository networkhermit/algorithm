#include <fun/vac/util/RandomFactory.h>
#include <fun/vac/util/TestRunner.h>

bool testIntegerN(void) {
    RandomFactory_seed();

    int value;
    for (int i = 0; i < 8192; i++) {
        if (RandomFactory_integerN(0, 0) != 0) {
            return false;
        }

        if (RandomFactory_integerN(1, 1) != 1) {
            return false;
        }

        value = RandomFactory_integerN(0, 1);
        if (value < 0 || value > 1) {
            return false;
        }

        value = RandomFactory_integerN(100, 10000);
        if (RandomFactory_integerN(value, value) != value) {
            return false;
        }
        if (value < 100 || value > 10000) {
            return false;
        }
    }

    return true;
}

bool testGenerateEven(void) {
    RandomFactory_seed();

    for (int i = 0; i < 8192; i++) {
        if ((RandomFactory_generateEven() & 1) != 0) {
            return false;
        }
    }

    return true;
}

bool testGenerateOdd(void) {
    RandomFactory_seed();

    for (int i = 0; i < 8192; i++) {
        if ((RandomFactory_generateOdd() & 1) == 0) {
            return false;
        }
    }

    return true;
}

int main(void) {
    TestRunner_parseTest(testIntegerN());

    TestRunner_parseTest(testGenerateEven());

    TestRunner_parseTest(testGenerateOdd());
}
