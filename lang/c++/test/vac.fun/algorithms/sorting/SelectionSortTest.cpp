#include <vac.fun/algorithms/sorting/SelectionSort.hpp>
#include <vac.fun/util/SequenceBuilder.hpp>
#include <vac.fun/util/Sequences.hpp>
#include <vac.fun/util/TestRunner.hpp>

using namespace std;

bool testSelectionSort() {
    size_t size = 32768;

    vector<int> arr(size);
    SequenceBuilder::packRandom(arr);

    int checksum = Sequences::parityChecksum(arr);

    SelectionSort::sort(arr);

    if (Sequences::parityChecksum(arr) != checksum) {
        return false;
    }

    return Sequences::isSorted(arr);
}

int main() {
    TestRunner::parseTest(testSelectionSort());
}
