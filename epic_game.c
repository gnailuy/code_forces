#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define NUM_PLAYER 2

int gcd(int x, int y) {
    if (0 == x) {
        return y;
    } else if (0 == y) {
        return x;
    } else {
        int r;
        do {
            r = x % y;
            x = y;
            y = r;
        } while (r != 0);
        return x;
    }
}

int main (int argc, char * argv[]) {
    int player[NUM_PLAYER], n, i = 0, n_taken;
    scanf ("%d %d %d", &player[0], &player[1], &n);

    while (true) {
        n_taken = gcd(player[i], n);
        i = (i + 1) % 2;

        if (n >= n_taken) {
            n -= n_taken;
        } else {
            break;
        }
    }
    printf ("%d\n", i);

    return 0;
}

