#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

void swap (int * buffer, int i, int j) {
    buffer[i] ^= buffer[j];
    buffer[j] ^= buffer[i];
    buffer[i] ^= buffer[j];
}

void bubble_sort (int * buffer, int length) {
    for (int i = 0; i < length; i++) {
        for (int j = i + 1; j < length; j++) {
            if (buffer[i] > buffer[j]) {
                swap (buffer, i, j);
            }
        }
    }
}

int main (int argc, char * argv[]) {
    int number_columns;
    scanf ("%d", &number_columns);

    int * columns = malloc (sizeof(int) * number_columns);
    for (int i = 0; i < number_columns; i++) {
        scanf ("%d", &columns[i]);
    }

    bubble_sort (columns, number_columns);

    for (int i = 0; i < number_columns; i++) {
        printf ("%d ", columns[i]);
    }
    printf ("\n");

    free (columns);
    return 0;
}

