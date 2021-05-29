#include <vac.fun/util/TestRunner.hpp>

using namespace std;

void testParseTest() {
    for (int i = 0; i < 10; i++) {
        TestRunner::parseTest((i & 1) != 0);
    }
}

int main() {
    testParseTest();
}
