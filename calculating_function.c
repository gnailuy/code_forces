#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    long long n;
    scanf ("%lld", &n);
    // scanf ("%I64d", &n);

    long long f = 0;
    if (0 == n % 2) {
        f = n / 2;
    } else {
        f = (n - 1) / 2 - n;
    }

    printf ("%lld\n", f);
    // printf ("%I64d\n", f);

    return 0;
}

