#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define TABLE_SIZE 100

int main (int argc, char * argv[]) {
    int n, x1, y1, x2, y2;
    scanf ("%d", &n);

    int * table[TABLE_SIZE] = { NULL };
    for (int i = 0; i < TABLE_SIZE; i++) {
        table[i] = (int *) calloc (TABLE_SIZE, sizeof(int)); // Zero-initialized
    }

    int range_x[2] = { TABLE_SIZE, -1 };
    int range_y[2] = { TABLE_SIZE, -1 };
    for (int k = 0; k < n; k++) {
        scanf ("%d %d %d %d", &x1, &y1, &x2, &y2);
        if (range_x[0] > x1 - 1) range_x[0] = x1 - 1;
        if (range_x[1] < x2 - 1) range_x[1] = x2 - 1;
        if (range_y[0] > y1 - 1) range_y[0] = y1 - 1;
        if (range_y[1] < y2 - 1) range_y[1] = y2 - 1;

        for (int i = x1 - 1; i <= x2 - 1; i++) {
            for (int j = y1 - 1; j <= y2 - 1; j++) {
                table[i][j]++;
            }
        }
    }

    int sum = 0;
    for (int i = range_x[0]; i <= range_x[1]; i++) {
        for (int j = range_y[0]; j <= range_y[1]; j++) {
            sum += table[i][j];
        }
    }
    printf ("%d\n", sum);

    for (int i = 0; i < TABLE_SIZE; i++) {
        free (table[i]);
    }
    return 0;
}

