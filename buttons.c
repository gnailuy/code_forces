#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, result = 0;
    scanf ("%d", &n);

    for (int i = n; i > 0; i--) {
        result += (i - 1) * (n - i + 1) + 1;
    }

    printf ("%d\n", result);

    return 0;
}

