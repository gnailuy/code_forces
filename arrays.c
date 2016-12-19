#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int na, nb, k, m;
    scanf ("%d %d", &na, &nb);
    scanf ("%d %d", &k, &m);

    int n, max_a, min_b;
    for (int i = 0; i < na; i++) {
        scanf ("%d", &n);
        if (i + 1 == k) {
            max_a = n;
        }
    }
    for (int i = 0; i < nb; i++) {
        scanf ("%d", &n);
        if (i == nb - m) {
            min_b = n;
        }
    }

    if (max_a < min_b) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    return 0;
}
