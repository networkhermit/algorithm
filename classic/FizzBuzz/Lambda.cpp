#include <iostream>

using namespace std;

string fizzBuzz(int n) {
    string word;

    if (n % 3 == 0) {
        word += "Fizz";
    }
    if (n % 5 == 0) {
        word += "Buzz";
    }

    if (word.empty()) {
        word = to_string(n);
    }

    return word;
}

int main() {
    for (int i = 1; i <= 100; i++) {
        cout << fizzBuzz(i) << endl;
    }
}
