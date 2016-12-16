#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, k;
    scanf ("%d %d\n", &n, &k);

    int * walks = (int *) malloc (sizeof(int) * (n + 1));
    walks[0] = k;

    int p, addition = 0;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &p);
        if (p + walks[i] >= k) {
            walks[i + 1] = p;
        } else {
            walks[i + 1] = k - walks[i];
            addition += walks[i + 1] - p;
        }
    }

    printf ("%d\n", addition);
    for (int i = 1; i < n; i++) {
        printf ("%d ", walks[i]);
    }
    printf ("%d\n", walks[n]);

    free (walks);
    return 0;
}
