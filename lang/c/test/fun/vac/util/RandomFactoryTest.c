#include <fun/vac/util/RandomFactory.h>
#include <fun/vac/util/TestRunner.h>

bool testGenerateEven(void) {
    RandomFactory_launch();

    int value;
    for (int i = 0; i < 8192; i++) {
        value = RandomFactory_generateEven();
        if ((value & 1) != 0) {
            return false;
        }
    }

    return true;
}

bool testGenerateOdd(void) {
    RandomFactory_launch();

    int value;
    for (int i = 0; i < 8192; i++) {
        value = RandomFactory_generateOdd();
        if ((value & 1) == 0) {
            return false;
        }
    }

    return true;
}

int main(void) {

    TestRunner_parseTest(testGenerateEven());

    TestRunner_parseTest(testGenerateOdd());
}
