#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, m;
    scanf ("%d %d", &n, &m);

    int * bulbs = (int *) calloc (m, sizeof(int)); // Zero-initialized

    int turns, lamp;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &turns);
        for (int j = 0; j < turns; j++) {
            scanf ("%d", &lamp);
            bulbs[lamp - 1] = 1;
        }
    }

    bool has_zero = false;
    for (int i = 0; i < m; i++) {
        if (0 == bulbs[i]) {
            has_zero = true;
            break;
        }
    }

    if (has_zero) {
        printf ("NO\n");
    } else {
        printf ("YES\n");
    }

    free (bulbs);
    return 0;
}

