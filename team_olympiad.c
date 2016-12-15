#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define GIFTS 3

int min (int * a, int length) {
    int _min = a[0];
    for (int i = 1; i < length; i++) {
        if (_min > a[i]) {
            _min = a[i];
        }
    }
    return _min;
}

int main (int argc, char * argv[]) {
    int n, gift;
    scanf ("%d", &n);

    if (n < GIFTS) {
        printf ("0\n");
        return 0;
    }

    int ** gifted_students = (int **) malloc (sizeof(int *) * GIFTS);
    for (int i = 0; i < GIFTS; i++) {
        gifted_students[i] = (int *) malloc (sizeof(int) * n);
    }

    int count[GIFTS] = { 0 };
    for (int i = 0; i < n; i++) {
        scanf ("%d", &gift);
        gifted_students[gift - 1][count[gift - 1]++] = i + 1;
    }

    int min_gift = min (count, GIFTS);
    printf ("%d\n", min_gift);
    for (int i = 0; i < min_gift; i++) {
        for (int j = 0; j < GIFTS; j++) {
            printf ("%d", gifted_students[j][i]);
            if (j != GIFTS - 1) printf (" ");
        }
        printf ("\n");
    }

    for (int i = 0; i < GIFTS; i++) {
        free (gifted_students[i]);
    }
    free (gifted_students);
    return 0;
}

