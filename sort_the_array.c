#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    if (n == 1) {
        scanf ("%d", &n); // Discard the input
        printf ("yes\n");
        printf ("1 1\n");
        return 0;
    }

    int * array = (int *) malloc (sizeof(int) * n); // Distinct
    for (int i = 0; i < n; i++) {
        scanf ("%d", &array[i]);
    }

    int ds_start = 0, ds_end = 0;
    bool can_sort = true;
    for (int i = 1; i < n; i++) {
        if (array[i - 1] > array[i]) {
            if (ds_end != 0) {
                can_sort = false;
                break;
            }
            if (ds_start == 0) ds_start = i;
        } else if (ds_start != 0) {
            if (ds_end == 0) ds_end = i;
        }
    }
    if (ds_end == 0) ds_end = n;

    if (can_sort && ds_start != 0) {
        if ((ds_end != n && array[ds_start - 1] > array[ds_end])
            || (ds_start != 1 && array[ds_end - 1] < array[ds_start - 2])) {
            can_sort = false;
        }
    }

    if (can_sort && ds_start != 0) {
        printf ("yes\n");
        printf ("%d %d\n", ds_start, ds_end);
    } else if (can_sort && ds_start == 0) {
        printf ("yes\n");
        printf ("1 1\n");
    } else {
        printf ("no\n");
    }

    free (array);
    return 0;
}

