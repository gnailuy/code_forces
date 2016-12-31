#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define GOOD_CAR 5

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int ** matrix = (int **) malloc (sizeof(int *) * n);
    for (int i = 0; i < n; i++) {
        matrix[i] = (int *) malloc (sizeof(int) * n);
        for (int j = 0; j < n; j++) {
            scanf ("%d", &matrix[i][j]);
        }
    }

    int good_car = 0;
    for (int i = 0; i < n; i++) {
        bool is_good_car = true;
        for (int j = 0; j < n; j++) {
            if (matrix[i][j] == 1 || matrix[i][j] == 3) {
                is_good_car = false;
                break;
            }
        }
        if (is_good_car) {
            matrix[i][0] = GOOD_CAR;
            good_car++;
        }
    }
    printf ("%d\n", good_car);
    if (good_car > 0) {
        for (int i = 0; i < n; i++) {
            if (GOOD_CAR == matrix[i][0]) {
                printf ("%d ", i + 1);
            }
        }
        printf ("\n");
    }

    for (int i = 0; i < n; i++) {
        free (matrix[i]);
    }
    free (matrix);

    return 0;
}

