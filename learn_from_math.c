#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define PANIC -1

bool is_composite (int n) {
    double upper = n / 2.0;
    for (int i = 2; i <= upper; i++) {
        if (n % i == 0) {
            return true;
        }
    }
    return false;
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    double upper = n / 2.0;
    for (int i = 4; i <= upper; i++) {
        if (is_composite (i)) {
            if (is_composite (n - i)) {
                printf ("%d %d\n", i, n - i);
                return 0;
            }
        }
    }

    printf ("According to the math, this output is impossible.\n");
    return PANIC;
}

