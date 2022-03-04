#include <cctype>
#include <iostream>
#include <stack>
#include <string>

using namespace std;

void process(string coordinate) {
    int column = 0;
    int row = 0;

    int length = coordinate.size();

    int k = 1;
    if (coordinate[0] == 'R' && !isalpha(coordinate[k])) {
        while (k < length && isdigit(coordinate[k])) {
            k++;
        }
    }

    if (k == 1 || k == length) {
        k = 0;
        while (isalpha(coordinate[k])) {
            column = column * 26 + (static_cast<int>(coordinate[k]) - 64);
            k++;
        }
        while (k < length) {
            row = row * 10 + (static_cast<int>(coordinate[k]) - 48);
            k++;
        }

        cout << "R" << row << "C" << column << endl;
    } else {
        k = 1;
        while (isdigit(coordinate[k])) {
            row = row * 10 + (static_cast<int>(coordinate[k]) - 48);
            k++;
        }
        k++;
        while (k < length) {
            column = column * 10 + (static_cast<int>(coordinate[k]) - 48);
            k++;
        }

        stack<char> s;

        int r;
        while (column != 0) {
            r = column % 26;
            column /= 26;
            if (r == 0) {
                s.push('Z');
                column--;
            } else {
                s.push(static_cast<char>(r + 64));
            }
        }

        while (!s.empty()) {
            cout << s.top();
            s.pop();
        }
        cout << row << endl;
    }
}

int main() {
    int testCases;

    string coordinate;

    while (cin >> testCases) {
        for (int i = 0; i < testCases; i++) {
            cin >> coordinate;
            process(coordinate);
        }
    }
}
