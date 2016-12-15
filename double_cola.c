#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define NUM_PEOPLE 5

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int i = 0, j = 0;
    int i_length = 0, step = 1;
    while (true) {
        i_length = NUM_PEOPLE * ((2 << i) - 1); // 5 * (2^i - 1)
        if (i_length >= n) {
            break;
        }
        i++;
    }
    if (i > 0) {
        i_length = NUM_PEOPLE * ((2 << (i - 1)) - 1);
        step = 2 << (i - 1);
    } else {
        i_length = 0;
        step = 1;
    }

    for (j = 0; j < NUM_PEOPLE; j++) {
        if (i_length + step * (j + 1) >= n) {
            break;
        }
    }

    if (0 == j) {
        printf ("Sheldon\n");
    } else if (1 == j) {
        printf ("Leonard\n");
    } else if (2 == j) {
        printf ("Penny\n");
    } else if (3 == j) {
        printf ("Rajesh\n");
    } else if (4 == j) {
        printf ("Howard\n");
    }

    return 0;
}

