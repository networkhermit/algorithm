#include <fun/vac/util/RandomFactory.hpp>
#include <fun/vac/util/TestRunner.hpp>

using namespace std;

bool testGenerateEven() {
    RandomFactory::launch();

    int value;
    for (int i = 0; i < 8192; i++) {
        value = RandomFactory::generateEven();
        if ((value & 1) != 0) {
            return false;
        }
    }

    return true;
}

bool testGenerateOdd() {
    RandomFactory::launch();

    int value;
    for (int i = 0; i < 8192; i++) {
        value = RandomFactory::generateOdd();
        if ((value & 1) == 0) {
            return false;
        }
    }

    return true;
}

int main() {

    TestRunner::parseTest(testGenerateEven());

    TestRunner::parseTest(testGenerateOdd());
}
