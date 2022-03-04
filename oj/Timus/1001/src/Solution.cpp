#include <cmath>
#include <iomanip>
#include <iostream>
#include <stack>

using namespace std;

int main() {
    stack<long long> s;

    long long n;

    while (cin >> n) {
        s.push(n);
    }

    while (!s.empty()) {
        cout << fixed << setprecision(4) << sqrt(static_cast<double>(s.top())) << endl;
        s.pop();
    }
}
