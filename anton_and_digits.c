#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int min (int n, ...) {
    va_list ks;
    va_start (ks, n);

    int res = va_arg (ks, int), arg;
    for (int i = 1; i < n; i++) {
        arg = va_arg (ks, int);
        if (arg < res) {
            res = arg;
        }
    }
    va_end (ks);

    return res;
}

int main (int argc, char * argv[]) {
    int k2, k3, k5, k6;
    scanf ("%d %d %d %d", &k2, &k3, &k5, &k6);

    int num256 = min (3, k2, k5, k6);
    k2 -= num256;
    int num32 = min (2, k2, k3);

    printf ("%d\n", 256 * num256 + 32 * num32);

    return 0;
}

