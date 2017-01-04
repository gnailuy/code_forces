#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, k;
    scanf ("%d %d", &n, &k);

    int f, t, joy, max_joy;
    scanf ("%d %d", &f, &t);
    max_joy = t > k ? f + k - t : f;
    for (int i = 1; i < n; i++) {
        scanf ("%d %d", &f, &t);
        joy = t > k ? f + k - t : f;
        if (joy > max_joy) max_joy = joy;
    }

    printf ("%d\n", max_joy);

    return 0;
}

