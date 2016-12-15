#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, size1, size2;
    scanf ("%d", &n);

    if (n < 2) {
        scanf ("%d", &size2);
        if (0 == size2) {
            printf ("UP\n");
        } else if (15 == size2) {
            printf ("DOWN\n");
        } else {
            printf ("-1\n");
        }

        return 0;
    }

    for (int i = 0; i < n - 1; i++) {
        scanf ("%d", &size1);
    }
    scanf ("%d", &size2);

    if (0 == size2) {
        printf ("UP\n");
    } else if (15 == size2) {
        printf ("DOWN\n");
    } else if (size1 < size2) {
        printf ("UP\n");
    } else {
        printf ("DOWN\n");
    }

    return 0;
}

