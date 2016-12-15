#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_LENGTH 1000

int main (int argc, char * argv[]) {
    int n, length;
    scanf ("%d", &n);

    int * table = (int *) calloc (MAX_LENGTH, sizeof(int)); // Zero-initialized
    for (int i = 0; i < n; i++) {
        scanf ("%d", &length);
        table[length - 1]++;
    }

    int tower_count = 0, max_height = 0;
    for (int i = 0; i < MAX_LENGTH; i++) {
        if (table[i] > 0) {
            tower_count++;
            if (max_height < table[i]) {
                max_height = table[i];
            }
        }
    }

    printf ("%d %d\n", max_height, tower_count);

    return 0;
}

