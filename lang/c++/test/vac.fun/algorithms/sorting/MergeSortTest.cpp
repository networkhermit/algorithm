#include <vac.fun/algorithms/sorting/MergeSort.hpp>
#include <vac.fun/util/SequenceBuilder.hpp>
#include <vac.fun/util/Sequences.hpp>
#include <vac.fun/util/TestRunner.hpp>

using namespace std;

bool testMergeSort() {
    size_t size = 32768;

    vector<int> arr(size);
    SequenceBuilder::packRandom(arr);

    int checksum = Sequences::parityChecksum(arr);

    MergeSort::sort(arr);

    if (Sequences::parityChecksum(arr) != checksum) {
        return false;
    }

    return Sequences::isSorted(arr);
}

int main() {
    TestRunner::parseTest(testMergeSort());
}
