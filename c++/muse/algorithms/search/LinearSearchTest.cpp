#include <muse/algorithms/search/LinearSearch.hpp>
#include <muse/util/SequenceBuilder.hpp>
#include <muse/util/TestRunner.hpp>

using namespace std;

bool testLinearSearch() {
    size_t size = 32768;

    vector<int> arr(size);
    SequenceBuilder::packIncreasing(arr);

    if (LinearSearch::find(arr, -1) != size) {
        return false;
    }

    if (LinearSearch::find(arr, 2'147'483'647) != size) {
        return false;
    }

    for (size_t i = 0; i < size; i++) {
        if (LinearSearch::find(arr, arr[i]) != i) {
            return false;
        }
    }

    return true;
}

int main() {
    TestRunner::pick(&testLinearSearch);
}
