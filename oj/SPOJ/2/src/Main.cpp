#include <cmath>
#include <cstring>
#include <iostream>

using namespace std;

void process(int m, int n) {
    if (n == 1) {
        return;
    }

    if (m <= 2) {
        cout << 2 << endl;
        m = 3;
    }

    if (n == 2) {
        return;
    }

    m |= 1;
    n = (n - 1) | 1;

    int cache = static_cast<int>(sqrt(static_cast<double>(n)));

    bool sievedFactors[(cache + 1) >> 1];
    memset(sievedFactors, false, (cache + 1) >> 1);

    for (int i = 3, bound = cache + 1; i < bound; i += 2) {
        if (sievedFactors[i >> 1]) {
            continue;
        }
        for (int j = i * i; j < bound; j += i << 1) {
            if (!sievedFactors[j >> 1]) {
                sievedFactors[j >> 1] = true;
            }
        }
    }

    bool nums[((n - m) >> 1) + 1];
    memset(nums, false, ((n - m) >> 1) + 1);

    int factor;
    int target;

    bool boost = m <= 9;

    for (int i = 1, length = sizeof(sievedFactors) / sizeof(sievedFactors[0]); i < length; i++) {
        if (sievedFactors[i]) {
            continue;
        }

        factor = (i << 1) + 1;

        if (boost) {
            target = factor * factor;
        } else if (m <= factor * factor) {
            boost = true;
            target = factor * factor;
        } else {
            if (m % factor == 0) {
                target = m;
            } else {
                target = m + (factor - m % factor);
                if ((target & 1) == 0) {
                    target += factor;
                }
            }
        }

        for (int j = target; j <= n; j += factor << 1) {
            if (!nums[(j - m) >> 1]) {
                nums[(j - m) >> 1] = true;
            }
        }
    }

    for (int i = 0, length = *(&nums + 1) - nums; i < length; i++) {
        if (!nums[i]) {
            cout << m + (i << 1) << endl;
        }
    }
}

int main() {
    int testCases;

    int m;
    int n;

    while (cin >> testCases) {
        for (int i = 0; i < testCases; i++) {
            cin >> m >> n;
            if (i != 0) {
                cout << endl;
            }
            process(m, n);
        }
    }
}
