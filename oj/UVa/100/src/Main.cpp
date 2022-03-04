#include <iostream>

using namespace std;

int process(int left, int right) {
    int n;

    int cycleLength;

    int maxCycleLength = 0;

    for (int i = left; i <= right; i++) {
        cycleLength = 0;

        n = i;

        while (true) {
            cycleLength++;
            if (n == 1) {
                if (cycleLength > maxCycleLength) {
                    maxCycleLength = cycleLength;
                }
                break;
            }
            if ((n & 1) == 0) {
                n >>= 1;
            } else {
                n = (n << 1) + n + 1;
            }
        }
    }

    return maxCycleLength;
}

int main() {
    int m;
    int n;

    while (cin >> m >> n) {
        if (m < n) {
            cout << m << " " << n << " " << process(m, n) << endl;
        } else {
            cout << m << " " << n << " " << process(n, m) << endl;
        }
    }
}
