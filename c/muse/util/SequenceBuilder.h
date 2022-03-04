#ifndef MUSE_UTIL_SEQUENCE_BUILDER_H
#define MUSE_UTIL_SEQUENCE_BUILDER_H

#include <stddef.h>

#include <muse/util/RandomFactory.h>
#include <muse/util/Sequences.h>

void SequenceBuilder_packIncreasing(int *arr, size_t length) {
    RandomFactory_seed();
    arr[0] = RandomFactory_genIntN(1, 3);
    for (size_t i = 1; i < length; i++) {
        arr[i] = arr[i - 1] + RandomFactory_genIntN(1, 3);
    }
}

void SequenceBuilder_packRandom(int *arr, size_t length) {
    RandomFactory_seed();
    for (size_t i = 0; i < length; i++) {
        arr[i] = RandomFactory_genInt();
    }
}

void SequenceBuilder_packDecreasing(int *arr, size_t length) {
    SequenceBuilder_packIncreasing(arr, length);
    Sequences_reverse(arr, length);
}

#endif
