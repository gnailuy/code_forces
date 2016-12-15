#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

// Find the index i where piles[i - 1] < target <= piles[i]
// Existence guaranteed
int find_pile (int * piles, int start, int end, int target) {
    if (end - start <= 1) {
        return end;
    } else {
        int middle = (end + start) / 2;
        if (piles[middle] >= target) {
            return find_pile (piles, start, middle, target);
        } else {
            return find_pile (piles, middle, end, target);
        }
    }
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int * piles = (int *) malloc (sizeof(int) * (n + 1));
    piles[0] = 0;
    scanf ("%d", &piles[1]);
    for (int i = 2; i < n + 1; i++) {
        scanf ("%d", &piles[i]);
        piles[i] += piles[i - 1];
    }

    int m, juicy;
    scanf ("%d", &m);
    for (int i = 0; i < m; i++) {
        scanf ("%d", &juicy);
        printf ("%d\n", find_pile (piles, 0, n, juicy));
    }

    free (piles);
    return 0;
}

