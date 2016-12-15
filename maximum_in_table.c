#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int ** array = (int **) malloc (sizeof(int *) * n);
    for (int i = 0; i < n; i++) {
        array[i] = (int *) malloc (sizeof(int) * n);
    }

    for (int i = 0; i < n; i++) {
        array[0][i] = 1;
        array[i][0] = 1;
    }
    for (int i = 1; i < n; i++) {
        for (int j = 1; j < n; j++) {
            array[i][j] = array[i - 1][j] + array[i][j - 1];
        }
    }

    printf ("%d\n", array[n - 1][n - 1]);

    for (int i = 0; i < n; i++) {
        free (array[i]);
    }
    free (array);
    return 0;
}

