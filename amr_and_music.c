#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

void swap (int * buffer, int i, int j) {
    buffer[i] ^= buffer[j];
    buffer[j] ^= buffer[i];
    buffer[i] ^= buffer[j];
}

void bubble_sort (int * buffer, int * index, int length) {
    for (int i = 0; i < length; i++) {
        for (int j = i + 1; j < length; j++) {
            if (buffer[i] > buffer[j]) {
                swap (buffer, i, j);
                swap (index, i, j);
            }
        }
    }
}

int main (int argc, char * argv[]) {
    int n, k;
    scanf ("%d %d", &n, &k);

    int * instruments = (int *) malloc (sizeof(int) * n);
    int * index = (int *) malloc (sizeof(int) * n);
    for (int i = 0; i < n; i++) {
        scanf ("%d", &instruments[i]);
        index[i] = i + 1;
    }
    bubble_sort (instruments, index, n);

    int to_learn = 0, cost_days = 0;
    for (int i = 0; i < n; i++) {
        if (cost_days + instruments[i] > k) {
            break;
        }
        to_learn++;
        cost_days += instruments[i];
    }

    printf ("%d\n", to_learn);
    for (int i = 0; i < to_learn; i++) {
        printf ("%d ", index[i]);
    }
    if (to_learn > 0) {
        printf ("\n");
    }

    free (instruments);
    free (index);
    return 0;
}

