#include <muse/algorithms/number_theory/FibonacciNumber.h>
#include <muse/util/TestRunner.h>

bool testFibonacciNumber(void) {
  long sample[][2] = {
      {-31, 1'346'269}, {-30, -832040}, {-29, 514229},   {-28, -317811},
      {-27, 196418},    {-26, -121393}, {-25, 75025},    {-24, -46368},
      {-23, 28657},     {-22, -17711},  {-21, 10946},    {-20, -6765},
      {-19, 4181},      {-18, -2584},   {-17, 1597},     {-16, -987},
      {-15, 610},       {-14, -377},    {-13, 233},      {-12, -144},
      {-11, 89},        {-10, -55},     {-9, 34},        {-8, -21},
      {-7, 13},         {-6, -8},       {-5, 5},         {-4, -3},
      {-3, 2},          {-2, -1},       {-1, 1},         {0, 0},
      {1, 1},           {2, 1},         {3, 2},          {4, 3},
      {5, 5},           {6, 8},         {7, 13},         {8, 21},
      {9, 34},          {10, 55},       {11, 89},        {12, 144},
      {13, 233},        {14, 377},      {15, 610},       {16, 987},
      {17, 1597},       {18, 2584},     {19, 4181},      {20, 6765},
      {21, 10946},      {22, 17711},    {23, 28657},     {24, 46368},
      {25, 75025},      {26, 121393},   {27, 196418},    {28, 317811},
      {29, 514229},     {30, 832040},   {31, 1'346'269},
  };

  for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
    if (FibonacciNumber_iterativeProcedure(sample[i][0]) != sample[i][1]) {
      return false;
    }
  }

  for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
    if (FibonacciNumber_recursiveProcedure(sample[i][0]) != sample[i][1]) {
      return false;
    }
  }

  return true;
}

int main(void) { TestRunner_pick(&testFibonacciNumber); }
