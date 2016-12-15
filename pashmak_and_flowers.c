#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_BEAUTY 1000000000 // 10^9

int main (int argc, char * argv[]) {
    int n, min = MAX_BEAUTY + 1, max = 0;
    scanf ("%d", &n);

    int * beauties = (int *) malloc (sizeof(int) * n);
    for (int i = 0; i < n; i++) {
        scanf ("%d", &beauties[i]);
        if (beauties[i] < min) {
            min = beauties[i];
        }
        if (beauties[i] > max) {
            max = beauties[i];
        }
    }

    int min_count = 0, max_count = 0;
    long long combination;
    if (max == min) {
        combination = ((long long) n) * ((long long) (n - 1)) / 2;
    } else {
        for (int i = 0; i < n; i++) {
            if (beauties[i] == min) {
                min_count++;
            } else if (beauties[i] == max) {
                max_count++;
            }
        }
        combination = ((long long) min_count) * ((long long) max_count);
    }

    // printf ("%d %I64d\n", max - min, combination);
    printf ("%d %lld\n", max - min, combination);

    free (beauties);
    return 0;
}

