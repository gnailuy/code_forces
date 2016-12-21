#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int i = 1;
    long long numbers = 0, nn = 9, digits = 0;
    while (numbers + nn <= n) {
        numbers += nn;
        digits += i * nn;
        nn *= 10; i++;
    }
    digits += i * (n - numbers);

    // printf ("%I64d\n", digits);
    printf ("%lld\n", digits);

    return 0;
}

