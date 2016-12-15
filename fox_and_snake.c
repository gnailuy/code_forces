#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, m;
    scanf ("%d %d", &n, &m);

    for (int i = 1; i < n + 1; i++) {
        if (i % 2 != 0) {
            for (int j = 0; j < m; j++) {
                printf ("#");
            }
        } else if ((i / 2) % 2 != 0) {
            for (int j = 0; j < m - 1; j++) {
                printf (".");
            }
            printf ("#");
        } else {
            printf ("#");
            for (int j = 0; j < m - 1; j++) {
                printf (".");
            }
        }
        printf ("\n");
    }

    return 0;
}

