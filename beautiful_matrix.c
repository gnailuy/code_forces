#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAT_SIZE 5 // Odd number
#define CENTER (MAT_SIZE / 2)

int abs (int n) {
    if (n >= 0) {
        return n;
    } else {
        return -n;
    }
}

int main (int argc, char * argv[]) {
    int x = CENTER, y = CENTER, in;

    for (int i = 0; i < MAT_SIZE; i++) {
        for (int j = 0; j < MAT_SIZE; j++) {
            scanf ("%d", &in);
            if (1 == in) {
                x = i; y = j;
            }
        }
    }

    int result = abs (x - CENTER) + abs (y - CENTER);
    printf ("%d\n", result);

    return 0;
}

