#include <stdio.h>
#include <stdlib.h>

#define NUM_FRIENDS 3

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
    int pos[NUM_FRIENDS];
    for (int i = 0; i < NUM_FRIENDS; i++) {
        scanf ("%d", &pos[i]);
    }

    bubble_sort (pos, NUM_FRIENDS);

    printf ("%d\n", pos[2] - pos[0]);

    return 0;
}

