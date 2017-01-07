#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int result = 0;
    if (n % 2 == 0) {
        int half = n / 2;
        if (half % 2 == 0) {
            result = half / 2 - 1;
        } else {
            result = half / 2;
        }
    }

    printf ("%d\n", result);

    return 0;
}

