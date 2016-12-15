#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MASK 0x00000001

int main (int argc, char * argv[]) {
    int n, bit, count = 0;
    scanf ("%d", &n);

    while (n > 0) {
        bit = n & MASK;
        n = n >> 1;
        if (bit == 1) {
            count++;
        }
    }

    printf ("%d\n", count);

    return 0;
}

