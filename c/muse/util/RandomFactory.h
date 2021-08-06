#ifndef MUSE_UTIL_RANDOM_FACTORY_H
#define MUSE_UTIL_RANDOM_FACTORY_H 1

#include <stdlib.h>
#include <time.h>

void RandomFactory_seed(void) {
    srand((unsigned) time(NULL));
}

int RandomFactory_genIntN(int min, int max) {
    return min + rand() % (max - min + 1);
}

int RandomFactory_genInt(void) {
    return RandomFactory_genIntN(1, 2147483647);
}

int RandomFactory_genEven(void) {
    return RandomFactory_genInt() >> 1 << 1;
}

int RandomFactory_genOdd(void) {
    return RandomFactory_genInt() | 1;
}

#endif
