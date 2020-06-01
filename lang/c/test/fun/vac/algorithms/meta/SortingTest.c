#include <fun/vac/algorithms/meta/Sorting.h>
#include <fun/vac/util/SequenceBuilder.h>
#include <fun/vac/util/Sequences.h>
#include <fun/vac/util/TestRunner.h>

bool testBubbleSort(void) {
    size_t size = 32768;

    int arr[size];
    SequenceBuilder_packRandom(arr, size);

    int checksum = Sequences_parityChecksum(arr, size);

    Sorting_bubbleSort(arr, size);

    if (Sequences_parityChecksum(arr, size) != checksum) {
        return false;
    }

    return Sequences_isSorted(arr, size);
}

bool testInsertionSort(void) {
    size_t size = 32768;

    int arr[size];
    SequenceBuilder_packRandom(arr, size);

    int checksum = Sequences_parityChecksum(arr, size);

    Sorting_insertionSort(arr, size);

    if (Sequences_parityChecksum(arr, size) != checksum) {
        return false;
    }

    return Sequences_isSorted(arr, size);
}

bool testMergeSort(void) {
    size_t size = 32768;

    int arr[size];
    SequenceBuilder_packRandom(arr, size);

    int checksum = Sequences_parityChecksum(arr, size);

    Sorting_mergeSort(arr, size);

    if (Sequences_parityChecksum(arr, size) != checksum) {
        return false;
    }

    return Sequences_isSorted(arr, size);
}

bool testQuickSort(void) {
    size_t size = 32768;

    int arr[size];
    SequenceBuilder_packRandom(arr, size);

    int checksum = Sequences_parityChecksum(arr, size);

    Sorting_quickSort(arr, size);

    if (Sequences_parityChecksum(arr, size) != checksum) {
        return false;
    }

    return Sequences_isSorted(arr, size);
}

bool testSelectionSort(void) {
    size_t size = 32768;

    int arr[size];
    SequenceBuilder_packRandom(arr, size);

    int checksum = Sequences_parityChecksum(arr, size);

    Sorting_selectionSort(arr, size);

    if (Sequences_parityChecksum(arr, size) != checksum) {
        return false;
    }

    return Sequences_isSorted(arr, size);
}

int main(void) {
    TestRunner_parseTest(testBubbleSort());

    TestRunner_parseTest(testInsertionSort());

    TestRunner_parseTest(testMergeSort());

    TestRunner_parseTest(testQuickSort());

    TestRunner_parseTest(testSelectionSort());
}
