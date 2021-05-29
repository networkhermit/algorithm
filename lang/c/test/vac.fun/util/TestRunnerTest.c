#include <vac.fun/util/TestRunner.h>

void testParseTest(void) {
    for (int i = 0; i < 10; i++) {
        TestRunner_parseTest((i & 1) != 0);
    }
}

int main(void) {
    testParseTest();
}
