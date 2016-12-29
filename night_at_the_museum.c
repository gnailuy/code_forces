#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define CHARS 26

int abs (int n) {
    return n >= 0 ? n : -n;
}

int main (int argc, char * argv[]) {
    char c = ' ', last_c = 'a';

    int moves = 0;
    while (true) {
        scanf ("%c", &c);
        if (c == '\n') break;

        int distance = abs (c - last_c);
        if (distance > CHARS / 2) {
            distance = CHARS - distance;
        }
        moves += distance;

        last_c = c;
    }

    printf ("%d\n", moves);

    return 0;
}

