#include <muse/util/TestRunner.hpp>

using namespace std;

void testPick() {
  for (int i = 0; i < 10; i++) {
    TestRunner::pick([i]() -> bool { return (i & 1) != 0; });
  }
}

int main() { testPick(); }
