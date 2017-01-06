#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int max;
    if (n % 2 == 0) {
        max = n * n / 2;
    } else {
        max = (n - 1) * (n - 1) / 2 + n;
    }
    printf ("%d\n", max);

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            if (i % 2 == 0) {
                if (j % 2 == 0) {
                    printf ("C");
                } else {
                    printf (".");
                }
            } else {
                if (j % 2 != 0) {
                    printf ("C");
                } else {
                    printf (".");
                }
            }
        }
        printf ("\n");
    }

    return 0;
}
