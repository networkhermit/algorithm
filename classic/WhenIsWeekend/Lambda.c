#include <stdio.h>
#include <time.h>

int main(void) {
    puts("When’s Weekend?");

    time_t t = time(NULL);
    int today = localtime(&t)->tm_wday;

    switch (today) {
    case 0:
        puts("Today.");
        break;
    case 6:
        puts("Today.");
        break;
    case 5:
        puts("Tomorrow.");
        break;
    case 4:
        puts("In two days.");
        break;
    default:
        puts("Too far away.");
    }
}
