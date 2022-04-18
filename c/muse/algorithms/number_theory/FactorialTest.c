#include <muse/algorithms/number_theory/Factorial.h>
#include <muse/util/TestRunner.h>

bool testFactorial(void) {
  long sample[][2] = {
      {0, 1},        {1, 1},         {2, 2},          {3, 6},     {4, 24},
      {5, 120},      {6, 720},       {7, 5040},       {8, 40320}, {9, 362880},
      {10, 3628800}, {11, 39916800}, {12, 479001600},
  };

  for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
    if (Factorial_iterativeProcedure(sample[i][0]) != sample[i][1]) {
      return false;
    }
  }

  for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
    if (Factorial_recursiveProcedure(sample[i][0]) != sample[i][1]) {
      return false;
    }
  }

  return true;
}

int main(void) { TestRunner_pick(&testFactorial); }
