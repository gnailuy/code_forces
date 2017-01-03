#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    if (n < 10) {
        printf ("%d\n", n);
    } else if (n < 190) {
        if (n % 2 == 0) {
            printf ("%d\n", (n - 9) / 20 + 1);
        } else {
            printf ("%d\n", ((n - 9) / 2 - 1) % 10);
        }
    } else {
        if (n % 3 == 1) {
            printf ("%d\n", ((n - 190) / 3) / 100 + 1);
        } else if (n % 3 == 2) {
            printf ("%d\n", (((n - 190 - 1) / 3) / 10) % 10);
        } else {
            printf ("%d\n", ((n - 190 - 2) / 3) % 10);
        }
    }

    return 0;
}

