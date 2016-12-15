#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int ** intersections = (int **) malloc (sizeof(int *) * n);
    for (int i = 0; i < n; i++) {
        intersections[i] = (int *) calloc (n, sizeof(int)); // Zero-initialized
    }

    int x, y;
    for (int i = 0; i < n * n; i++) {
        scanf ("%d %d", &x, &y);
        x -= 1; y -= 1;
        if (intersections[x][y] == 0) { // Make progress
            printf ("%d ", i + 1);
            for (int j = 0; j < n; j++) {
                intersections[x][j] = 1;
                intersections[j][y] = 1;
            }
        }
    }
    printf ("\n");

    for (int i = 0; i < n; i++) {
        free (intersections[i]);
    }
    free (intersections);
    return 0;
}

