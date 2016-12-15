#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int number_soldiers = 0;
    scanf ("%d", &number_soldiers);

    int * heights = (int *) malloc (sizeof(int) * number_soldiers);
    for (int i = 0; i < number_soldiers; i++) {
        scanf ("%d", &heights[i]);
    }

    int max = 0, min = 101;
    int first_max_position, last_min_position;
    for (int i = 0; i < number_soldiers; i++) {
        if (heights[i] > max) {
            max = heights[i];
            first_max_position = i;
        }
        if (heights[i] <= min) {
            min = heights[i];
            last_min_position = i;
        }
    }
    int across = 0;
    if (first_max_position > last_min_position) {
        across = 1;
    }

    printf ("%d\n",
        first_max_position + number_soldiers - 1 - last_min_position - across);

    free (heights);
    return 0;
}

