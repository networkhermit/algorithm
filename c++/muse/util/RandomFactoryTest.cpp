#include <muse/util/RandomFactory.hpp>
#include <muse/util/TestRunner.hpp>

using namespace std;

bool testGenIntN() {
    RandomFactory::seed();

    int value;
    for (int i = 0; i < 8192; i++) {
        if (RandomFactory::genIntN(0, 0) != 0) {
            return false;
        }

        if (RandomFactory::genIntN(1, 1) != 1) {
            return false;
        }

        value = RandomFactory::genIntN(0, 1);
        if (value < 0 || value > 1) {
            return false;
        }

        value = RandomFactory::genIntN(100, 10000);
        if (RandomFactory::genIntN(value, value) != value) {
            return false;
        }
        if (value < 100 || value > 10000) {
            return false;
        }
    }

    return true;
}

bool testGenEven() {
    RandomFactory::seed();

    for (int i = 0; i < 8192; i++) {
        if ((RandomFactory::genEven() & 1) != 0) {
            return false;
        }
    }

    return true;
}

bool testGenOdd() {
    RandomFactory::seed();

    for (int i = 0; i < 8192; i++) {
        if ((RandomFactory::genOdd() & 1) == 0) {
            return false;
        }
    }

    return true;
}

int main() {
    TestRunner::pick(&testGenIntN);

    TestRunner::pick(&testGenEven);

    TestRunner::pick(&testGenOdd);
}
