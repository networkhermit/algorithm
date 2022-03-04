#include <algorithm>
#include <cmath>
#include <iostream>
#include <vector>

using namespace std;

vector<bool> bruteForce(int n) {
    vector<bool> arr(n);

    fill(arr.begin(), arr.end(), false);

    for (int i = 1; i <= n; i++) {
        for (int j = i - 1; j < n; j += i) {
            arr[j] = !arr[j];
        }
    }

    return arr;
}

vector<bool> process(int n) {
    vector<bool> arr(n);

    int temp;
    for (int i = 1; i <= n; i++) {
        temp = static_cast<int>(sqrt(static_cast<double>(i)));
        arr[i - 1] = temp * temp == i;
    }

    return arr;
}

int main() {
    const int N = 100;

    vector<bool> arr = process(N);

    for (int i = 0; i < N; i++) {
        if (arr[i]) {
            cout << i + 1 << "\t  OPEN" << endl;
        }
    }
}
