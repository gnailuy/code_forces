#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int * seq = (int *) malloc (sizeof(int) * n);
    for (int i = 0; i < n; i++) {
        scanf ("%d", &seq[i]);
    }

    int * max = (int *) malloc (sizeof(int) * n);
    max[0] = 1 - seq[0];
    int max_turn = max[0], ones = seq[0];

    for (int i = 1; i < n; i++) {
        if (seq[i] == 0) {
            if (max[i - 1] > 0) {
                max[i] = max[i - 1] + 1;
            } else {
                max[i] = 1;
            }
        } else { // seq[i] == 1
            ones += 1;
            if (max[i - 1] > 0) {
                max[i] = max[i - 1] - 1;
            } else {
                max[i] = 0;
            }
        }
        if (max_turn < max[i]) {
            max_turn = max[i];
        }
    }

    if (max_turn > 0) {
        printf ("%d\n", max_turn + ones);
    } else {
        printf ("%d\n", ones - 1);
    }

    free (seq);
    free (max);
    return 0;
}

