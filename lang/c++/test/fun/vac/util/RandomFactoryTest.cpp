#include <fun/vac/util/RandomFactory.hpp>
#include <fun/vac/util/TestRunner.hpp>

using namespace std;

bool testIntegerN() {
    RandomFactory::seed();

    int value;
    for (int i = 0; i < 8192; i++) {
        if (RandomFactory::integerN(0, 0) != 0) {
            return false;
        }

        if (RandomFactory::integerN(1, 1) != 1) {
            return false;
        }

        value = RandomFactory::integerN(0, 1);
        if (value < 0 || value > 1) {
            return false;
        }

        value = RandomFactory::integerN(100, 10000);
        if (RandomFactory::integerN(value, value) != value) {
            return false;
        }
        if (value < 100 || value > 10000) {
            return false;
        }
    }

    return true;
}

bool testGenerateEven() {
    RandomFactory::seed();

    for (int i = 0; i < 8192; i++) {
        if ((RandomFactory::generateEven() & 1) != 0) {
            return false;
        }
    }

    return true;
}

bool testGenerateOdd() {
    RandomFactory::seed();

    for (int i = 0; i < 8192; i++) {
        if ((RandomFactory::generateOdd() & 1) == 0) {
            return false;
        }
    }

    return true;
}

int main() {

    TestRunner::parseTest(testIntegerN());

    TestRunner::parseTest(testGenerateEven());

    TestRunner::parseTest(testGenerateOdd());
}
