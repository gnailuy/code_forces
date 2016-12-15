#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

void insertion_sort (int * array, int length) {
    int j, t;
    for (int i = 1; i < length; i++) {
        j = i;
        while (j > 0 && array[j - 1] > array[j]) {
            t = array[j];
            array[j] = array[j - 1];
            array[j - 1] = t;
            j--;
        }
    }
}

double max (double a, double b) {
    return a > b ? a : b;
}

int main (int argc, char * argv[]) {
    int n, l;
    scanf ("%d %d", &n, &l);

    int * pos = malloc (sizeof(int) * n);
    for (int i = 0; i < n; i++) {
        scanf ("%d", &pos[i]);
    }

    insertion_sort (pos, n);

    double dist_head = pos[0] - 0;
    double dist_tail = l - pos[n - 1];

    double d, dist_middle = 0;
    for (int i = 1; i < n; i++) {
        d = (pos[i] - pos[i - 1]) / 2.0;
        if (d > dist_middle) {
            dist_middle = d;
        }
    }

    printf ("%0.10f\n", max (dist_head, max (dist_middle, dist_tail)));

    free (pos);
    return 0;
}

