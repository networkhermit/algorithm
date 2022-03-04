#include <iostream>

using namespace std;

long long process(int n, int m, int a) {
    int x = (n - 1) / a + 1;
    int y = (m - 1) / a + 1;

    return static_cast<long long>(x) * static_cast<long long>(y);
}

int main() {
    int n;
    int m;
    int a;

    while (cin >> n >> m >> a) {
        cout << process(n, m, a) << endl;
    }
}
