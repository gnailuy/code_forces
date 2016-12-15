#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    long n, k, result;
    scanf ("%ld %ld", &n, &k);
    // long long n, k, result;
    // scanf ("%I64d %I64d", &n, &k);

    long delimiter = n - n / 2;
    // long long delimiter = n - n / 2;
    if (k <= delimiter) {
        result = 2 * k - 1;
    } else {
        result = 2 * (k - delimiter);
    }

    printf ("%ld\n", result);
    // printf ("%I64d\n", result);

    return 0;
}

