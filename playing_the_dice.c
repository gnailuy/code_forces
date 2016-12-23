#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int abs (int n) {
    return n >= 0 ? n : -n;
}

int main (int argc, char * argv[]) {
    int a, b;
    scanf ("%d %d", &a, &b);

    int aw = 0, draw = 0, bw = 0;
    for (int i = 1; i <= 6; i++) {
        int ad = abs (a - i);
        int bd = abs (b - i);
        if (ad < bd) {
            aw++;
        } else if (ad == bd) {
            draw++;
        } else {
            bw++;
        }
    }

    printf ("%d %d %d\n", aw, draw, bw);

    return 0;
}

