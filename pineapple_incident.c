#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int t, s, x;
    scanf ("%d %d %d", &t, &s, &x);

    if (x == t || (x - t > 1 && ((x - t) % s == 0 || (x - t) % s == 1))) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    return 0;
}

