#include <vac.fun/algorithms/search/BinarySearch.hpp>
#include <vac.fun/util/SequenceBuilder.hpp>
#include <vac.fun/util/TestRunner.hpp>

using namespace std;

bool testBinarySearch() {
    size_t size = 32768;

    vector<int> arr(size);
    SequenceBuilder::packIncreasing(arr);

    if (BinarySearch::find(arr, -1) != size) {
        return false;
    }

    if (BinarySearch::find(arr, 2'147'483'647) != size) {
        return false;
    }

    for (size_t i = 0; i < size; i++) {
        if (BinarySearch::find(arr, arr[i]) != i) {
            return false;
        }
    }

    return true;
}

int main() {
    TestRunner::parseTest(testBinarySearch());
}
