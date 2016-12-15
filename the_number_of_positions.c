#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, a, b;
    scanf ("%d %d %d", &n, &a, &b);

    printf ("%d\n", n + 1 - (a + 1 > n - b ? a + 1 : n - b));

    return 0;
}

