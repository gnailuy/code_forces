#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int m, s;
    scanf ("%d %d", &m, &s);

    if (s < 0 || s > m * 9) {
        printf ("-1 -1\n");
        return 0;
    } else if (0 == s) {
        if (1 == m) {
            printf ("0 0\n");
        } else {
            printf ("-1 -1\n");
        }
        return 0;
    }

    int * min = (int *) calloc (m, sizeof(int)); // Zero-initialized
    int * max = (int *) calloc (m, sizeof(int)); // Zero-initialized
    min[0] = max[0] = 1;

    int min_pos = m - 1;
    int max_pos = 0;
    for (int i = 2; i <= s; i++) {
        min[min_pos]++;
        if (min[min_pos] == 9) {
            min_pos--;
        }

        max[max_pos]++;
        if (max[max_pos] == 9) {
            max_pos++;
        }
    }

    for (int i = 0; i < m; i++) {
        printf ("%d", min[i]);
    }
    printf (" ");
    for (int i = 0; i < m; i++) {
        printf ("%d", max[i]);
    }
    printf ("\n");

    free (min);
    free (max);
    return 0;
}

