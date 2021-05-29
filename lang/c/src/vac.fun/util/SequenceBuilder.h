#ifndef VAC_FUN_UTIL_SEQUENCE_BUILDER_H
#define VAC_FUN_UTIL_SEQUENCE_BUILDER_H 1

#include <stddef.h>

#include <vac.fun/util/RandomFactory.h>
#include <vac.fun/util/Sequences.h>

void SequenceBuilder_packIncreasing(int *arr, size_t length) {
    RandomFactory_seed();
    arr[0] = RandomFactory_integerN(1, 3);
    for (size_t i = 1; i < length; i++) {
        arr[i] = arr[i - 1] + RandomFactory_integerN(1, 3);
    }
}

void SequenceBuilder_packRandom(int *arr, size_t length) {
    RandomFactory_seed();
    for (size_t i = 0; i < length; i++) {
        arr[i] = RandomFactory_generateInteger();
    }
}

void SequenceBuilder_packDecreasing(int *arr, size_t length) {
    SequenceBuilder_packIncreasing(arr, length);
    Sequences_reverse(arr, length);
}

#endif
