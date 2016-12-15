#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int teams, result = 0;
    scanf ("%d", &teams);

    int * home = (int *) malloc (sizeof(int) * teams);
    int * guest = (int *) malloc (sizeof(int) * teams);

    for (int i = 0; i < teams; i++) {
        scanf ("%d %d", &home[i], &guest[i]);
    }

    for (int i = 0; i < teams; i++) {
        for (int j = 0; j < teams; j++) {
            if (i != j && home[i] == guest[j]) {
                result += 1;
            }
        }
    }

    printf ("%d\n", result);

    free (home);
    free (guest);
    return 0;
}

