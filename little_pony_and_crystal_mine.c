#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n; // n is odd
    scanf ("%d", &n);

    int half = n / 2;
    for (int i = 0; i < n; i++) {
        int d, s;
        if (i <= half) {
            d = 1 + 2 * i;
            s = (n - d) / 2;
        } else {
            d = 1 + 2 * (n - i - 1);
            s = (n - d) / 2;
        }

        for (int j = 0; j < s; j++) {
            printf ("*");
        }
        for (int j = 0; j < d; j++) {
            printf ("D");
        }
        for (int j = 0; j < s; j++) {
            printf ("*");
        }
        printf ("\n");
    }

    return 0;
}

