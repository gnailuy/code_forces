#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int numbers = 0, i = 1, nn = 9, digits = 0;
    while (numbers + nn <= n) {
        numbers += nn;
        digits += i * nn;
        nn *= 10; i++;
    }
    digits += i * (n - numbers);

    printf ("%d\n", digits);

    return 0;
}

