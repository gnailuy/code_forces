#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MONTHS 12

void swap (int * buffer, int i, int j) {
    buffer[i] ^= buffer[j];
    buffer[j] ^= buffer[i];
    buffer[i] ^= buffer[j];
}

void bubble_sort (int * buffer, int length) {
    for (int i = 0; i < length; i++) {
        for (int j = i + 1; j < length; j++) {
            if (buffer[i] < buffer[j]) { // Descending
                swap (buffer, i, j);
            }
        }
    }
}

int main (int argc, char * argv[]) {
    int target, max = 0, current = 0, work_month = 0;
    scanf ("%d", &target);

    int * grows = (int *) malloc (sizeof(int) * MONTHS);
    for (int i = 0; i < MONTHS; i++) {
        scanf ("%d", &grows[i]);
        max += grows[i];
    }
    if (target > max) {
        printf ("-1\n");
        free (grows);
        return 0;
    }
    bubble_sort (grows, MONTHS);
    for (int i = 0; i < MONTHS; i++) {
        if (target <= current) {
            break;
        }
        current += grows[i];
        work_month++;
    }
    printf ("%d\n", work_month);

    free (grows);
    return 0;
}

