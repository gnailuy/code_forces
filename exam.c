#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    if (n <= 2) {
        printf ("1\n");
        printf ("1\n");
    } else if (3 == n) {
        printf ("2\n");
        printf ("1 3\n");
    } else {
        printf ("%d\n", n);
        if (n % 2 != 0) {
            printf ("%d ", n--);
        }
        int half = n / 2;
        for (int i = 1; i <= half; i++) {
            printf ("%d %d ", half + i, i);
        }
        printf ("\n");
    }

    return 0;
}

