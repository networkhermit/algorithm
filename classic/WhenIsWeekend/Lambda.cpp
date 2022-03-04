#include <ctime>
#include <iostream>

using namespace std;

int main() {
    cout << "When’s Weekend?" << endl;

    time_t t = time(nullptr);
    int today = localtime(&t)->tm_wday;

    switch (today) {
    case 0:
        cout << "Today." << endl;
        break;
    case 6:
        cout << "Today." << endl;
        break;
    case 5:
        cout << "Tomorrow." << endl;
        break;
    case 4:
        cout << "In two days." << endl;
        break;
    default:
        cout << "Too far away." << endl;
    }
}
