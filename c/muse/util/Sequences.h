#ifndef MUSE_UTIL_SEQUENCES_H
#define MUSE_UTIL_SEQUENCES_H

#include <stdbool.h>
#include <stdio.h>

#include <muse/algorithms/sorting/QuickSort.h>
#include <muse/util/RandomFactory.h>

void Sequences_inspect(const int *arr, size_t length) {
    puts("[");
    for (size_t i = 0; i < length; i++) {
        printf("\t#%04zX  -->  %d\n", i, arr[i]);
    }
    puts("]");
}

bool Sequences_isSorted(const int *arr, size_t length) {
    for (size_t i = 1; i < length; i++) {
        if (arr[i] < arr[i - 1]) {
            return false;
        }
    }

    return true;
}

int Sequences_parityChecksum(const int *arr, size_t length) {
    int checksum = 0;

    for (size_t i = 0; i < length; i++) {
        checksum ^= arr[i];
    }

    return checksum;
}

void Sequences_reverse(int *arr, size_t length) {
    int k;
    int temp;

    for (size_t i = 0, bound = length >> 1; i < bound; i++) {
        k = length - i - 1;
        temp = arr[i];
        arr[i] = arr[k];
        arr[k] = temp;
    }
}

void Sequences_shuffle(int *arr, size_t length) {
    int k;
    int temp;

    RandomFactory_seed();
    for (size_t i = 0; i < length; i++) {
        k = RandomFactory_genIntN(i, length);
        temp = arr[i];
        arr[i] = arr[k];
        arr[k] = temp;
    }
}

void Sequences_sort(int *arr, size_t length) {
    QuickSort_sort(arr, length);
}

#endif
