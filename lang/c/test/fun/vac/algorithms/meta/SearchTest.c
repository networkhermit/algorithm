#include <fun/vac/algorithms/meta/Search.h>
#include <fun/vac/util/SequenceBuilder.h>
#include <fun/vac/util/TestRunner.h>

bool testBinarySearch(void) {
    size_t size = 32768;

    int arr[size];
    SequenceBuilder_packIncreasing(arr, size);

    if (Search_binarySearch(arr, size, -1) != size) {
        return false;
    }

    if (Search_binarySearch(arr, size, 2147483647) != size) {
        return false;
    }

    for (size_t i = 0; i < size; i++) {
        if (Search_binarySearch(arr, size, arr[i]) != i) {
            return false;
        }
    }

    return true;
}

bool testLinearSearch(void) {
    size_t size = 32768;

    int arr[size];
    SequenceBuilder_packIncreasing(arr, size);

    if (Search_linearSearch(arr, size, -1) != size) {
        return false;
    }

    if (Search_linearSearch(arr, size, 2147483647) != size) {
        return false;
    }

    for (size_t i = 0; i < size; i++) {
        if (Search_linearSearch(arr, size, arr[i]) != i) {
            return false;
        }
    }

    return true;
}

int main(void) {
    TestRunner_parseTest(testBinarySearch());

    TestRunner_parseTest(testLinearSearch());
}
