#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int max (int a, int b) {
    return a > b ? a : b;
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int a, pos[2], pos_idx = 0;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &a);
        if (1 == a) {
            pos[pos_idx++] = i + 1;
        }
        if (n == a) {
            pos[pos_idx++] = i + 1;
        }
    }

    printf ("%d\n", max (pos[1] - 1, n - pos[0]));

    return 0;
}
