#include <fun/vac/algorithms/search/BinarySearch.h>
#include <fun/vac/util/SequenceBuilder.h>
#include <fun/vac/util/TestRunner.h>

bool testBinarySearch(void) {
    size_t size = 32768;

    int arr[size];
    SequenceBuilder_packIncreasing(arr, size);

    if (BinarySearch_find(arr, size, -1) != size) {
        return false;
    }

    if (BinarySearch_find(arr, size, 2147483647) != size) {
        return false;
    }

    for (size_t i = 0; i < size; i++) {
        if (BinarySearch_find(arr, size, arr[i]) != i) {
            return false;
        }
    }

    return true;
}

int main(void) {
    TestRunner_parseTest(testBinarySearch());
}
