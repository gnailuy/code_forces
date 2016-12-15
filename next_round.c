#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, k, i;
    int participants = 0;

    scanf ("%d", &n);
    scanf ("%d", &k);

    int * seq = malloc (sizeof(int) * n);
    for (i = 0; i < n; i++) {
        scanf ("%d", &seq[i]);
    }

    for (i = 0; i < k; i++) {
        if (seq[i] > 0) {
            participants++;
        } else {
            printf ("%d\n", participants);
            free (seq);
            return 0;
        }
    }

    for (i = k; i < n; i++) {
        if (seq[i] >= seq[k - 1]) {
            participants++;
        } else {
            break;
        }
    }

    printf ("%d\n", participants);
    free (seq);
    return 0;
}
