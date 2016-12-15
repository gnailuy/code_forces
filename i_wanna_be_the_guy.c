#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int number_levels, p, q, l;
    scanf ("%d", &number_levels);

    int * levels = (int *) calloc (number_levels, sizeof(int)); // Zeros

    scanf ("%d", &p);
    for (int i = 0; i < p; i++) {
        scanf ("%d", &l);
        levels[l - 1] = 1;
    }
    scanf ("%d", &q);
    for (int i = 0; i < q; i++) {
        scanf ("%d", &l);
        levels[l - 1] = 1;
    }

    bool can_pass = true;
    for (int i = 0; i < number_levels; i++) {
        if (levels[i] == 0) {
            can_pass = false;
            break;
        }
    }
    if (can_pass) {
        printf ("I become the guy.\n");
    } else {
        printf ("Oh, my keyboard!\n");
    }

    free (levels);
    return 0;
}

