#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define NUM_CURRENT_HORSESHOES 4

int main (int argc, char * argv[]) {
    long colors[NUM_CURRENT_HORSESHOES] = {0L};
    for (int i = 0; i < NUM_CURRENT_HORSESHOES; i++) {
        scanf ("%ld", &colors[i]);
    }

    long distinct_colors[NUM_CURRENT_HORSESHOES] = {0};
    int index = 0;
    distinct_colors[0] = colors[0];
    for (int i = 1; i < NUM_CURRENT_HORSESHOES; i++) {
        bool dup_color = false;
        for (int j = 0; j < index + 1; j++) {
            if (colors[i] == distinct_colors[j]) {
                dup_color = true;
                break;
            }
        }
        if (!dup_color) {
            distinct_colors[++index] = colors[i];
        }
    }

    printf ("%d\n", NUM_CURRENT_HORSESHOES - index - 1);

    return 0;
}

