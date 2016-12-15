#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    long long a, b;
    // scanf ("%I64d %I64d", &a, &b);
    scanf ("%lld %lld", &a, &b);

    long long square = 0, remainder;
    do {
        square += a / b;
        remainder = a % b;
        a = b;
        b = remainder;
    } while (remainder != 0);

    // printf ("%I64d\n", square);
    printf ("%lld\n", square);

    return 0;
}

