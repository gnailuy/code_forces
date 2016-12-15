#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_CUPS 5
#define MAX_MEDALS 10

int main (int argc, char * argv[]) {
    int t, cups = 0, medals = 0, n, shelives_needed = 0;
    for (int i = 0; i < 3; i++) {
        scanf ("%d", &t);
        cups += t;
    }
    for (int i = 0; i < 3; i++) {
        scanf ("%d", &t);
        medals += t;
    }
    scanf ("%d", &n);

    if (cups % MAX_CUPS != 0) {
        shelives_needed += cups / MAX_CUPS + 1;
    } else {
        shelives_needed += cups / MAX_CUPS;
    }
    if (medals % MAX_MEDALS != 0) {
        shelives_needed += medals / MAX_MEDALS + 1;
    } else {
        shelives_needed += medals / MAX_MEDALS;
    }

    if (shelives_needed > n) {
        printf ("NO\n");
    } else {
        printf ("YES\n");
    }

    return 0;
}

