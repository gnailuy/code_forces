#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int a, b, c;
    scanf ("%d", &a);
    scanf ("%d", &b);
    scanf ("%d", &c);

    int max = 0;

    int res[6] = {
        a + b + c,
        a + b * c,
        (a + b) * c,
        a * b + c,
        a * (b + c),
        a * b * c
    };

    for (int i = 0; i < 6; i++) {
        if (res[i] > max) {
            max = res[i];
        }
    }
    printf ("%d\n", max);

    return 0;
}

