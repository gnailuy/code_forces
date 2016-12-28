#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

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
    int n;
    scanf ("%d", &n);

    int * badges = (int *) malloc (sizeof(int) * n);
    for (int i = 0; i < n; i++) {
        scanf ("%d", &badges[i]);
    }
    bubble_sort (badges, n);

    int cost = 0;
    for (int i = 1; i < n; i++) {
        if (badges[i] <= badges[i - 1]) {
            cost += badges[i - 1] - badges[i] + 1;
            badges[i] = badges[i - 1] + 1;
        }
    }

    printf ("%d\n", cost);

    free (badges);
    return 0;
}

