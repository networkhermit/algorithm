#include <muse/util/RandomFactory.h>
#include <muse/util/TestRunner.h>

bool testGenIntN(void) {
    RandomFactory_seed();

    int value;
    for (int i = 0; i < 8192; i++) {
        if (RandomFactory_genIntN(0, 0) != 0) {
            return false;
        }

        if (RandomFactory_genIntN(1, 1) != 1) {
            return false;
        }

        value = RandomFactory_genIntN(0, 1);
        if (value < 0 || value > 1) {
            return false;
        }

        value = RandomFactory_genIntN(100, 10000);
        if (RandomFactory_genIntN(value, value) != value) {
            return false;
        }
        if (value < 100 || value > 10000) {
            return false;
        }
    }

    return true;
}

bool testGenEven(void) {
    RandomFactory_seed();

    for (int i = 0; i < 8192; i++) {
        if ((RandomFactory_genEven() & 1) != 0) {
            return false;
        }
    }

    return true;
}

bool testGenOdd(void) {
    RandomFactory_seed();

    for (int i = 0; i < 8192; i++) {
        if ((RandomFactory_genOdd() & 1) == 0) {
            return false;
        }
    }

    return true;
}

int main(void) {
    TestRunner_pick(&testGenIntN);

    TestRunner_pick(&testGenEven);

    TestRunner_pick(&testGenOdd);
}
