#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int m, n, s;
    scanf ("%d %d", &m, &n);

    if (m > n) {
        s = n;
    } else {
        s = m;
    }

    if (0 != s % 2) {
        printf ("Akshat\n");
    } else {
        printf ("Malvika\n");
    }

    return 0;
}

