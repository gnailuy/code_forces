#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int balance;
    scanf ("%d", &balance);

    if (balance >= 0) {
        printf ("%d\n", balance);
    } else {
        int option1 = balance / 10;
        int option2 = balance / 100 * 10 + balance - balance / 10 * 10;
        printf ("%d\n", option1 > option2 ? option1 : option2);
    }

    return 0;
}
