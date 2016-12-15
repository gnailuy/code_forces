#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_LAPTOPS 100000

int main (int argc, char * argv[]) {
    int n, price, quality;
    scanf ("%d", &n);

    int * qualities = (int *) calloc (MAX_LAPTOPS, sizeof(int)); // Zero-initialized
    for (int i = 0; i < n; i++) {
        scanf ("%d %d", &price, &quality);
        qualities[price - 1] = quality;
    }

    bool disordered = false;
    int current_max = 0;
    for (int i = 0; i < MAX_LAPTOPS; i++) {
        if (qualities[i] > 0) {
            if (qualities[i] < current_max) {
                disordered = true;
                break;
            }
            current_max = qualities[i];
        }
    }

    if (disordered) {
        printf ("Happy Alex\n");
    } else {
        printf ("Poor Alex\n");
    }

    free (qualities);
    return 0;
}

