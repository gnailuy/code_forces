#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, l;
    scanf ("%d", &n);

    int ls[4] = {6, 8, 4, 2};

    if (0 == n) {
        l = 1;
    } else {
        int m = n % 4;
        l = ls[m];
    }

    printf ("%d\n", l);

    return 0;
}

