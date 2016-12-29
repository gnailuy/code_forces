#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int max (int a, int b) {
    return a >= b ? a : b;
}

int main (int argc, char * argv[]) {
    int a, b, c, d;
    scanf ("%d %d %d %d", &a, &b, &c, &d);

    int misha = max (3 * a / 10, a - a / 250 * c);
    int vasya = max (3 * b / 10, b - b / 250 * d);

    if (misha > vasya) {
        printf ("Misha\n");
    } else if (misha < vasya) {
        printf ("Vasya\n");
    } else {
        printf ("Tie\n");
    }

    return 0;
}

