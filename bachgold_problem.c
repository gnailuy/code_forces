#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    if (n % 2 == 0) {
        printf ("%d\n", n / 2);
        for (int i = 0; i < n / 2; i++) {
            printf ("2 ");
        }
        printf ("\n");
    } else {
        printf ("%d\n", (n - 1) / 2);
        for (int i = 0; i < (n - 1) / 2 - 1; i++) {
            printf ("2 ");
        }
        printf ("3");
        printf ("\n");
    }

    return 0;
}

