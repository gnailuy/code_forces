#include <stdarg.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

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
    int n, k, l, c, d, p, nl, np;
    scanf ("%d %d %d %d %d %d %d %d", &n, &k, &l, &c, &d, &p, &nl, &np);

    printf ("%d\n", min (3, (k * l) / (nl * n), (c * d) / n, p / (np * n)));

    return 0;
}

