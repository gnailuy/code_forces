#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define STD_LENGTH 3

int main (int argc, char * argv[]) {
    int n, s[STD_LENGTH], s_min = -1, count[STD_LENGTH] = { 0 };
    scanf ("%d", &n);
    for (int i = 0; i < STD_LENGTH; i++) {
        scanf ("%d", &s[i]);
        if (s[i] < s_min && s_min != -1) {
            s_min = s[i];
        }
    }

    int * solutions = calloc (n + 1, sizeof(int)); // Zero-initialized
    for (int i = 0; i < STD_LENGTH; i++) {
        if (s[i] <= n) {
            solutions[s[i]] = 1;
        }
    }

    int possible;
    for (int i = s_min; i <= n; i++) {
        for (int j = 0; j < STD_LENGTH; j++) {
            for (int k = 1; k * s[j] <= i; k++) {
                if (solutions[i - k * s[j]] > 0) {
                    possible = solutions[i - k * s[j]] + k;
                    if (possible > solutions[i]) {
                        solutions[i] = possible;
                    }
                }
            }
        }
    }
    printf ("%d\n", solutions[n]);

    free (solutions);
    return 0;
}

