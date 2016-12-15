#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, m;
    scanf ("%d %d", &n, &m);

    int * scores = calloc (n, sizeof(int)); // Zero-initialized

    int votes, max, max_index;
    for (int i = 0; i < m; i++) { // m cities
        max = 0;
        max_index = 0;
        for (int j = 0; j < n; j++) {
            scanf ("%d", &votes);
            if (votes > max) {
                max_index = j;
                max = votes;
            }
        }
        scores[max_index]++;
    }

    max = 0; max_index = 0;
    for (int i = 0; i < n; i++) {
        if (scores[i] > max) {
            max_index = i;
            max = scores[i];
        }
    }

    printf ("%d\n", max_index + 1);

    free (scores);
    return 0;
}

