#include <muse/util/TestRunner.h>

bool test(void) {
  static int round = 0;
  bool result = (round & 1) != 0;
  round++;
  return result;
}

void testPick(void) {
  for (int i = 0; i < 10; i++) {
    TestRunner_pick(&test);
  }
}

int main(void) { testPick(); }
