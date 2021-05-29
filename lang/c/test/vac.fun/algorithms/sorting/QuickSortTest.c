#include <vac.fun/algorithms/sorting/QuickSort.h>
#include <vac.fun/util/SequenceBuilder.h>
#include <vac.fun/util/Sequences.h>
#include <vac.fun/util/TestRunner.h>

bool testQuickSort(void) {
    size_t size = 32768;

    int arr[size];
    SequenceBuilder_packRandom(arr, size);

    int checksum = Sequences_parityChecksum(arr, size);

    QuickSort_sort(arr, size);

    if (Sequences_parityChecksum(arr, size) != checksum) {
        return false;
    }

    return Sequences_isSorted(arr, size);
}

int main(void) {
    TestRunner_parseTest(testQuickSort());
}