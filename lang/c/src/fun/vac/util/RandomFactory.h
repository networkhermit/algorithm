#ifndef FUN_VAC_UTIL_RANDOM_FACTORY_H
#define FUN_VAC_UTIL_RANDOM_FACTORY_H 1

#include <stdlib.h>
#include <time.h>

void RandomFactory_seed(void) {
    srand((unsigned) time(NULL));
}

int RandomFactory_integerN(int min, int max) {
    return min + rand() % (max - min + 1);
}

int RandomFactory_generateInteger(void) {
    return RandomFactory_integerN(0, 2147483647);
}

int RandomFactory_generateEven(void) {
    return RandomFactory_generateInteger() >> 1 << 1;
}

int RandomFactory_generateOdd(void) {
    return RandomFactory_generateInteger() | 1;
}

#endif
