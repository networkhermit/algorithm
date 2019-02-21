#include <fun/vac/util/TestRunner.h>

void testParseTest(void) {
    for (int i = 0; i < 10; i++) {
        if ((i & 1) == 0) {
            TestRunner_parseTest(false);
        } else {
            TestRunner_parseTest(true);
        }
    }
}

int main(void) {
    testParseTest();
}
