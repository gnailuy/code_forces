#include <stdio.h>
#include <stdlib.h>

void swap (int * array, int i, int j) {
    if (i != j) {
        int tmp = array[i];
        array[i] = array[j];
        array[j] = tmp;
    }
}

int partition (int * array, int start, int end) {
    int pivot = array[end];
    int i = start;
    for (int j = start; j < end; j++) {
        if (array[j] <= pivot) {
            swap (array, i++, j);
        }
    }
    swap (array, i, end);
    return i;
}

void quick_sort (int * array, int start, int end) {
    if (start < end) {
        int p = partition (array, start, end);
        quick_sort (array, start, p - 1);
        quick_sort (array, p + 1, end);
    }
}

int main (int argc, char * argv[]) {
    int number_students, number_puzzles;
    scanf ("%d %d", &number_students, &number_puzzles);

    int * pieces = malloc (sizeof(int) * number_puzzles);
    for (int i = 0; i < number_puzzles; i++) {
        scanf ("%d", &pieces[i]);
    }

    quick_sort (pieces, 0, number_puzzles - 1);
    int result = -1, difference;
    for (int i = 0; i <= number_puzzles - number_students ; i++) {
        difference = pieces[i + number_students - 1] - pieces[i];
        if (difference < result || result < 0) {
            result = difference;
        }
    }
    printf ("%d\n", result);

    free (pieces);
    return 0;
}

