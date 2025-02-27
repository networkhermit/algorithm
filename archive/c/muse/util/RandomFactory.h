#ifndef MUSE_UTIL_RANDOM_FACTORY_H
#define MUSE_UTIL_RANDOM_FACTORY_H

#include <stdlib.h>
#include <time.h>

void RandomFactory_seed(void) { srand((unsigned)time(nullptr)); }

int RandomFactory_genIntN(int min, int max) {
  return min + rand() % (max - min + 1);
}

int RandomFactory_genInt(void) {
  return RandomFactory_genIntN(1, 2'147'483'647);
}

int RandomFactory_genEven(void) { return RandomFactory_genInt() >> 1 << 1; }

int RandomFactory_genOdd(void) { return RandomFactory_genInt() | 1; }

#endif
