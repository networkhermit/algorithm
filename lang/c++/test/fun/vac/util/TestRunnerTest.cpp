#include <fun/vac/util/TestRunner.hpp>

using namespace std;

void testParseTest() {
    for (int i = 0; i < 10; i++) {
        if ((i & 1) == 0) {
            TestRunner::parseTest(false);
        } else {
            TestRunner::parseTest(true);
        }
    }
}

int main() {
    testParseTest();
}
