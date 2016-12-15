#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int a, b, fashion, non_fashion;
    scanf ("%d %d", &a, &b);

    if (a > b) {
        int t = a;
        a = b;
        b = t;
    }

    fashion = a;
    non_fashion = (b - a) / 2;
    printf ("%d %d\n", fashion, non_fashion);

    return 0;
}

