#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, a, b;
    scanf ("%d %d %d", &n, &a, &b);

    int pos = b % n + a;

    if (pos > n) pos -= n;
    else if (pos <= 0) pos += n;

    printf ("%d\n", pos);

    return 0;
}

