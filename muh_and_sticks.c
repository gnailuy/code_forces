#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define NUM_STICKS 6
#define DIGITS 10

int main (int argc, char * argv[]) {
    int l, d[DIGITS] = { 0 };
    for (int i = 0; i < NUM_STICKS; i++) {
        scanf ("%d", &l);
        d[l]++;
    }

    bool has_leg = false;
    bool has_elephant_hb = false;
    for (int i = 1; i < DIGITS; i++) {
        if (d[i] >= 4) {
            has_leg = true;
            if (d[i] == 6) {
                has_elephant_hb = true;
            }
        } else if (d[i] == 2) {
            has_elephant_hb = true;
        }
    }

    if (has_leg) {
        if (has_elephant_hb) {
            printf ("Elephant\n");
        } else {
            printf ("Bear\n");
        }
    } else {
        printf ("Alien\n");
    }

    return 0;
}

