#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int min (int a, int b) {
    return a < b ? a : b;
}

int main (int argc, char * argv[]) {
    int d1, d2, d3, result;
    scanf ("%d %d %d", &d1, &d2, &d3);

    if (d1 < d2 + d3) {
        if (d1 + d3 < d2) {
            result = 2 * (d1 + d3);
        } else {
            result = min (2 * (d1 + d2), d1 + d2 + d3);
        }
    } else {
        result = 2 * (d2 + d3);
    }

    printf ("%d\n", result);

    return 0;
}

